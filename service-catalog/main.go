package main

import (
	"context"

	"service-catalog/internal"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {

	cfg := internal.LoadConfig()

	db := internal.ConnectDB(cfg.DatabaseURL)

	service := internal.NewCatalogService(db)

	handler := internal.NewCatalogHandler(service)

	app := fiber.New()

	api := app.Group(
		"/catalog",
	)

	// Packages
	api.Post("/packages", handler.CreatePackage)
	api.Get("/packages", handler.ListPackages)
	api.Get("/packages/:id", handler.GetPackage)
	api.Put("/packages/:id", handler.UpdatePackage)
	api.Delete("/packages/:id", handler.SoftDeletePackage)

	// Addons
	api.Post("/addons", handler.CreateAddon)
	api.Get("/addons", handler.ListAddons)
	api.Delete("/addons/:id", handler.SoftDeleteAddon)

	// Pricing Rules
	api.Post("/pricing-rules", handler.CreatePricingRule)
	api.Get("/pricing-rules", handler.ListPricingRules)

	// Price Calculation
	api.Post("/price/calculate", handler.CalculatePrice)

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
