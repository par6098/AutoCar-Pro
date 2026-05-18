package main

import (
	"context"

	"user-service/internal"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {

	cfg := internal.LoadConfig()

	db := internal.ConnectDB(cfg.DatabaseURL)

	service := internal.NewUserService(db)

	handler := internal.NewUserHandler(service)

	app := fiber.New()

	api := app.Group(
		"/users",
	)

	// Tenant
	api.Post("/tenants", handler.CreateTenant)
	api.Get("/tenants/:id/hierarchy", handler.GetTenantHierarchy)

	// Customers
	api.Post("/customers", handler.CreateCustomer)
	api.Get("/customers/:id", handler.GetCustomer)

	// Vehicles
	api.Post("/customers/:id/vehicles", handler.AddVehicle)
	api.Get("/customers/:id/vehicles", handler.ListVehicles)

	// Preferences
	api.Put("/customers/:id/preferences", handler.UpdatePreferences)

	// Employees
	api.Post("/employees", handler.CreateEmployee)
	api.Get("/employees/:id", handler.GetEmployee)
	api.Post("/employees/:id/skills", handler.AddEmployeeSkill)

	// Admins
	api.Post("/admins", handler.CreateAdmin)

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
