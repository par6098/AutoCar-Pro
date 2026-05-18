package main

import (
	"context"

	"booking-service/internal"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	cfg := internal.LoadConfig()

	db := internal.ConnectDB(cfg.DatabaseURL)
	redisClient := internal.ConnectRedis(cfg)

	service := internal.NewBookingService(db, redisClient, cfg)
	handler := internal.NewBookingHandler(service)

	app := fiber.New()

	api := app.Group(
		"/bookings",
	)

	api.Post("/", handler.CreateBooking)
	api.Get("/:id", handler.GetBooking)
	api.Put("/:id", handler.UpdateBooking)
	api.Post("/:id/cancel", handler.CancelBooking)
	api.Get("/slots/availability", handler.CheckSlotAvailability)

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
