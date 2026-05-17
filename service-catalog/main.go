package main

import (
	"log"

	"service-catalog/internal"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := internal.LoadConfig()

	db := internal.ConnectDB(cfg.DatabaseURL)
	defer db.Close()

	service := internal.NewCatalogService(db)
	handler := internal.NewCatalogHandler(service)

	app := fiber.New()

	api := app.Group("/catalog")

	api.Post("/packages", handler.CreatePackage)
	api.Get("/packages", handler.ListPackages)
	api.Get("/packages/:id", handler.GetPackage)
	api.Put("/packages/:id", handler.UpdatePackage)
	api.Delete("/packages/:id", handler.SoftDeletePackage)

	api.Post("/addons", handler.CreateAddon)
	api.Get("/addons", handler.ListAddons)
	api.Delete("/addons/:id", handler.SoftDeleteAddon)

	api.Post("/pricing-rules", handler.CreatePricingRule)
	api.Get("/pricing-rules", handler.ListPricingRules)
	api.Post("/price/calculate", handler.CalculatePrice)

	log.Fatal(app.Listen(":" + cfg.Port))
}
