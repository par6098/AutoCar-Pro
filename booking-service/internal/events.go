package internal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type EventPublisher struct {
	redisClient *redis.Client
	queueName   string
}

func NewEventPublisher(
	redisClient *redis.Client,
	queueName string,
) *EventPublisher {

	return &EventPublisher{
		redisClient: redisClient,
		queueName:   queueName,
	}
}

type BaseEvent struct {
	EventID       string      `json:"event_id"`
	EventType     string      `json:"event_type"`
	Source        string      `json:"source"`
	Version       string      `json:"version"`
	Timestamp     time.Time   `json:"timestamp"`
	CorrelationID string      `json:"correlation_id,omitempty"`
	Data          interface{} `json:"data"`
}

func (p *EventPublisher) Publish(
	ctx context.Context,
	eventType string,
	data interface{},
) error {

	event := BaseEvent{
		EventID:   generateEventID(),
		EventType: eventType,
		Source:    "booking-service",
		Version:   "1.0",
		Timestamp: time.Now(),
		Data:      data,
	}

	payload, err := json.Marshal(event)

	if err != nil {
		return err
	}

	err = p.redisClient.Publish(
		ctx,
		p.queueName,
		payload,
	).Err()

	if err != nil {
		return err
	}

	log.Printf(
		"event published: %s",
		eventType,
	)

	return nil
}

type EventSubscriber struct {
	redisClient *redis.Client
	queueName   string
}

func NewEventSubscriber(
	redisClient *redis.Client,
	queueName string,
) *EventSubscriber {

	return &EventSubscriber{
		redisClient: redisClient,
		queueName:   queueName,
	}
}

func (s *EventSubscriber) Subscribe(
	ctx context.Context,
	handler func(event BaseEvent),
) error {

	pubsub := s.redisClient.Subscribe(
		ctx,
		s.queueName,
	)

	defer pubsub.Close()

	ch := pubsub.Channel()

	log.Printf(
		"subscribed to queue: %s",
		s.queueName,
	)

	for msg := range ch {

		var event BaseEvent

		err := json.Unmarshal(
			[]byte(msg.Payload),
			&event,
		)

		if err != nil {
			log.Printf(
				"failed to unmarshal event: %v",
				err,
			)
			continue
		}

		handler(event)
	}

	return nil
}

func generateEventID() string {
	return time.Now().Format("20060102150405.000000")
}
