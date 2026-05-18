package internal

import (
	"context"
	"encoding/json"
	"log"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/redis/go-redis/v9"
)

func StartBookingEventConsumer(
	ctx context.Context,
	redisClient *redis.Client,
	ch clickhouse.Conn,
	cfg Config,
) {
	pubsub := redisClient.Subscribe(ctx, cfg.BookingEventQueue)
	defer pubsub.Close()

	service := NewAnalyticsService(ch)

	log.Println("analytics consumer subscribed to:", cfg.BookingEventQueue)

	for msg := range pubsub.Channel() {
		var event BookingEventEnvelope

		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			log.Println("failed to parse booking event:", err)
			continue
		}

		if event.EventType == "" {
			log.Println("invalid event received: missing event_type")
			continue
		}

		if err := service.InsertBookingEvent(ctx, event); err != nil {
			log.Println("failed to insert booking analytics event:", err)
			continue
		}

		log.Println("analytics event inserted:", event.EventType)
	}
}
