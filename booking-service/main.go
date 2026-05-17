package main

import (
	"log"

	"booking-service/internal"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := internal.LoadConfig()

	db := internal.ConnectDB(cfg.DatabaseURL)
	defer db.Close()

	redisClient := internal.ConnectRedis(cfg)
	defer redisClient.Close()

	bookingService := internal.NewBookingService(db, redisClient, cfg)
	handler := internal.NewBookingHandler(bookingService)

	app := fiber.New()

	api := app.Group("/bookings")

	api.Post("/", handler.CreateBooking)
	api.Get("/:id", handler.GetBooking)
	api.Put("/:id", handler.UpdateBooking)
	api.Post("/:id/cancel", handler.CancelBooking)
	api.Get("/slots/availability", handler.CheckSlotAvailability)

	log.Fatal(app.Listen(":" + cfg.Port))
}
