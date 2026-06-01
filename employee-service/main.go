package main

import (
	"context"

	"employee-service/internal"

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

	service := internal.NewEmployeeService(db)
	handler := internal.NewEmployeeHandler(service)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/employee")

	api.Post("/shifts", handler.CreateShift)
	api.Get("/shifts/:employee_id", handler.ListShifts)

	api.Post("/jobs/assign", handler.AssignJob)
	api.Put("/jobs/:id/status", handler.UpdateJobStatus)
	api.Get("/jobs/:employee_id", handler.ListJobs)

	api.Post("/attendance/check-in", handler.CheckIn)
	api.Post("/attendance/check-out", handler.CheckOut)

	api.Get("/performance/:employee_id", handler.GetPerformance)

	api.Post("/drivers/:employee_id/location", handler.UpdateDriverLocation)
	api.Get("/drivers/:employee_id/location", handler.GetDriverLocation)

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
