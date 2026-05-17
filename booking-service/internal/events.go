package internal

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type BookingEvent struct {
	EventID   string        `json:"event_id"`
	EventType string        `json:"event_type"`
	BookingID string        `json:"booking_id"`
	Status    BookingStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
}

func PublishBookingEvent(
	ctx context.Context,
	redisClient *redis.Client,
	queue string,
	event BookingEvent,
) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return redisClient.LPush(ctx, queue, payload).Err()
}
