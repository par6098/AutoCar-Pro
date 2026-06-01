package main

import (
	"context"

	"auth-service/internal"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {

	cfg := internal.LoadConfig()

	db := internal.ConnectDB(cfg.DatabaseURL)

	h := internal.NewHandler(db, cfg)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Public Routes
	app.Post("/auth/register", h.Register)
	app.Post("/auth/login", h.Login)
	app.Post("/auth/refresh", h.RefreshToken)

	// Protected Routes
	protected := app.Group("/api", internal.JWTMiddleware(cfg))

	protected.Get("/me", h.Me)

	// Admin
	protected.Get(
		"/admin/dashboard",
		internal.RequireRole("admin"),
		h.AdminDashboard,
	)

	// Employee
	protected.Get(
		"/employee/dashboard",
		internal.RequireRole("employee"),
		h.EmployeeDashboard,
	)

	// Customer
	protected.Get(
		"/customer/dashboard",
		internal.RequireRole("customer"),
		h.CustomerDashboard,
	)

	fiberLambda = fiberadapter.New(app)
}

func Handler(
	ctx context.Context,
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {

	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
