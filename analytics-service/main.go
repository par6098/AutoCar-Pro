package main

import (
	"context"

	"analytics-service/internal"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	cfg := internal.LoadConfig()

	ch := internal.ConnectClickHouse(cfg)
	redisClient := internal.ConnectRedis(cfg)

	service := internal.NewAnalyticsService(ch)
	handler := internal.NewAnalyticsHandler(service)

	go internal.StartBookingEventConsumer(context.Background(), redisClient, ch, cfg)

	app := fiber.New()

	api := app.Group("/analytics", internal.JWTMiddleware(cfg))

	api.Get("/revenue", handler.GetRevenue)
	api.Get("/employees/performance", handler.GetEmployeePerformance)
	api.Get("/services/popularity", handler.GetServicePopularity)
	api.Get("/bookings/summary", handler.GetBookingSummary)

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
