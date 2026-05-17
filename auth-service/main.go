package main

import (
	"log"

	"auth-service/internal"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := internal.LoadConfig()
	db := internal.ConnectDB(cfg.DatabaseURL)
	defer db.Close()

	app := fiber.New()

	h := internal.NewHandler(db, cfg)

	app.Post("/auth/register", h.Register)
	app.Post("/auth/login", h.Login)
	app.Post("/auth/refresh", h.RefreshToken)

	protected := app.Group("/api", internal.JWTMiddleware(cfg))
	protected.Get("/me", h.Me)

	admin := protected.Group("/admin", internal.RequireRole("admin"))
	admin.Get("/dashboard", h.AdminDashboard)

	employee := protected.Group("/employee", internal.RequireRole("employee"))
	employee.Get("/dashboard", h.EmployeeDashboard)

	customer := protected.Group("/customer", internal.RequireRole("customer"))
	customer.Get("/dashboard", h.CustomerDashboard)

	log.Fatal(app.Listen(":" + cfg.Port))
}