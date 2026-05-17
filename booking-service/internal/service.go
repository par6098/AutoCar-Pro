package internal

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type BookingService struct {
	db    *pgxpool.Pool
	redis *redis.Client
	cfg   Config
}

func NewBookingService(db *pgxpool.Pool, redisClient *redis.Client, cfg Config) *BookingService {
	return &BookingService{
		db:    db,
		redis: redisClient,
		cfg:   cfg,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req CreateBookingRequest) (*Booking, error) {
	available, err := s.IsSlotAvailable(ctx, req.ServiceID, req.SlotStart, req.SlotEnd)
	if err != nil {
		return nil, err
	}

	if !available {
		return nil, errors.New("slot not available")
	}

	bookingID := uuid.New().String()

	booking := &Booking{
		ID:            bookingID,
		CustomerID:    req.CustomerID,
		ServiceID:     req.ServiceID,
		SlotStart:     req.SlotStart,
		SlotEnd:       req.SlotEnd,
		PickupAddress: req.PickupAddress,
		DropAddress:   req.DropAddress,
		PickupTime:    req.PickupTime,
		DropTime:      req.DropTime,
		Status:        StatusCreated,
	}

	err = s.db.QueryRow(ctx,
		`INSERT INTO bookings 
		(id, customer_id, service_id, slot_start, slot_end, pickup_address, drop_address, pickup_time, drop_time, status, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW(),NOW())
		RETURNING created_at, updated_at`,
		booking.ID,
		booking.CustomerID,
		booking.ServiceID,
		booking.SlotStart,
		booking.SlotEnd,
		booking.PickupAddress,
		booking.DropAddress,
		booking.PickupTime,
		booking.DropTime,
		booking.Status,
	).Scan(&booking.CreatedAt, &booking.UpdatedAt)

	if err != nil {
		return nil, err
	}

	_ = s.emitEvent(ctx, booking.ID, "BOOKING_CREATED", booking.Status)

	return booking, nil
}

func (s *BookingService) GetBooking(ctx context.Context, id string) (*Booking, error) {
	var b Booking

	err := s.db.QueryRow(ctx,
		`SELECT id, customer_id, service_id, slot_start, slot_end, pickup_address, drop_address, pickup_time, drop_time, status, created_at, updated_at
		 FROM bookings WHERE id = $1`,
		id,
	).Scan(
		&b.ID,
		&b.CustomerID,
		&b.ServiceID,
		&b.SlotStart,
		&b.SlotEnd,
		&b.PickupAddress,
		&b.DropAddress,
		&b.PickupTime,
		&b.DropTime,
		&b.Status,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (s *BookingService) UpdateBooking(ctx context.Context, id string, req UpdateBookingRequest) (*Booking, error) {
	available, err := s.IsSlotAvailableForUpdate(ctx, id, req.SlotStart, req.SlotEnd)
	if err != nil {
		return nil, err
	}

	if !available {
		return nil, errors.New("slot not available")
	}

	var b Booking

	err = s.db.QueryRow(ctx,
		`UPDATE bookings
		 SET slot_start=$1, slot_end=$2, pickup_address=$3, drop_address=$4, pickup_time=$5, drop_time=$6, status=$7, updated_at=NOW()
		 WHERE id=$8 AND status != 'CANCELLED'
		 RETURNING id, customer_id, service_id, slot_start, slot_end, pickup_address, drop_address, pickup_time, drop_time, status, created_at, updated_at`,
		req.SlotStart,
		req.SlotEnd,
		req.PickupAddress,
		req.DropAddress,
		req.PickupTime,
		req.DropTime,
		StatusUpdated,
		id,
	).Scan(
		&b.ID,
		&b.CustomerID,
		&b.ServiceID,
		&b.SlotStart,
		&b.SlotEnd,
		&b.PickupAddress,
		&b.DropAddress,
		&b.PickupTime,
		&b.DropTime,
		&b.Status,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	_ = s.emitEvent(ctx, b.ID, "BOOKING_UPDATED", b.Status)

	return &b, nil
}

func (s *BookingService) CancelBooking(ctx context.Context, id string) error {
	_, err := s.db.Exec(ctx,
		`UPDATE bookings
		 SET status='CANCELLED', updated_at=NOW()
		 WHERE id=$1 AND status != 'CANCELLED'`,
		id,
	)

	if err != nil {
		return err
	}

	_ = s.emitEvent(ctx, id, "BOOKING_CANCELLED", StatusCancelled)

	return nil
}

func (s *BookingService) IsSlotAvailable(ctx context.Context, serviceID string, start, end time.Time) (bool, error) {
	var count int

	err := s.db.QueryRow(ctx,
		`SELECT COUNT(1)
		 FROM bookings
		 WHERE service_id = $1
		   AND status != 'CANCELLED'
		   AND slot_start < $3
		   AND slot_end > $2`,
		serviceID,
		start,
		end,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (s *BookingService) IsSlotAvailableForUpdate(ctx context.Context, bookingID string, start, end time.Time) (bool, error) {
	var count int

	err := s.db.QueryRow(ctx,
		`SELECT COUNT(1)
		 FROM bookings
		 WHERE id != $1
		   AND status != 'CANCELLED'
		   AND slot_start < $3
		   AND slot_end > $2`,
		bookingID,
		start,
		end,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (s *BookingService) emitEvent(ctx context.Context, bookingID, eventType string, status BookingStatus) error {
	return PublishBookingEvent(ctx, s.redis, s.cfg.BookingEventQueue, BookingEvent{
		EventID:   uuid.New().String(),
		EventType: eventType,
		BookingID: bookingID,
		Status:    status,
		CreatedAt: time.Now(),
	})
}
