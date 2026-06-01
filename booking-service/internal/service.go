package internal

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
)

type BookingService struct {
	db    *pgxpool.Pool
	redis *redis.Client
	cfg   Config
}

func (s *BookingService) GetBookingsByCustomer(ctx *fasthttp.RequestCtx, customerID string) (any, any) {
	panic("unimplemented")
}

func NewBookingService(db *pgxpool.Pool, redisClient *redis.Client, cfg Config) *BookingService {
	return &BookingService{
		db:    db,
		redis: redisClient,
		cfg:   cfg,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req CreateBookingRequest) (*Booking, error) {
	available, err := s.CheckSlotAvailability(ctx, req.BookingDate, req.BookingSlot)
	if err != nil {
		return nil, err
	}

	if !available {
		return nil, errors.New("slot not available")
	}

	id := uuid.New().String()
	var booking Booking

	err = s.db.QueryRow(ctx,
		`INSERT INTO bookings
		(id, customer_id, vehicle_id, service_package_id, booking_date, booking_slot,
		 pickup_required, pickup_address, status, created_at, updated_at)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW(),NOW())
		 RETURNING id, customer_id, vehicle_id, service_package_id, booking_date, booking_slot,
		 pickup_required, pickup_address, status, created_at, updated_at`,
		id,
		req.CustomerID,
		req.VehicleID,
		req.ServicePackageID,
		req.BookingDate,
		req.BookingSlot,
		req.PickupRequired,
		req.PickupAddress,
		BookingStatusCreated,
	).Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.VehicleID,
		&booking.ServicePackageID,
		&booking.BookingDate,
		&booking.BookingSlot,
		&booking.PickupRequired,
		&booking.PickupAddress,
		&booking.Status,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	_ = s.publishEvent(ctx, "BOOKING_CREATED", booking)

	return &booking, nil
}

func (s *BookingService) GetBooking(ctx context.Context, id string) (*Booking, error) {
	var booking Booking

	err := s.db.QueryRow(ctx,
		`SELECT id, customer_id, vehicle_id, service_package_id, booking_date, booking_slot,
		        pickup_required, pickup_address, status, created_at, updated_at
		 FROM bookings
		 WHERE id = $1`,
		id,
	).Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.VehicleID,
		&booking.ServicePackageID,
		&booking.BookingDate,
		&booking.BookingSlot,
		&booking.PickupRequired,
		&booking.PickupAddress,
		&booking.Status,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *BookingService) UpdateBooking(ctx context.Context, id string, req UpdateBookingRequest) (*Booking, error) {
	if req.Status == "" {
		req.Status = BookingStatusConfirmed
	}

	var booking Booking

	err := s.db.QueryRow(ctx,
		`UPDATE bookings
		 SET booking_date=$1,
		     booking_slot=$2,
		     pickup_required=$3,
		     pickup_address=$4,
		     status=$5,
		     updated_at=NOW()
		 WHERE id=$6
		   AND status != 'CANCELLED'
		 RETURNING id, customer_id, vehicle_id, service_package_id, booking_date, booking_slot,
		 pickup_required, pickup_address, status, created_at, updated_at`,
		req.BookingDate,
		req.BookingSlot,
		req.PickupRequired,
		req.PickupAddress,
		req.Status,
		id,
	).Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.VehicleID,
		&booking.ServicePackageID,
		&booking.BookingDate,
		&booking.BookingSlot,
		&booking.PickupRequired,
		&booking.PickupAddress,
		&booking.Status,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	_ = s.publishEvent(ctx, "BOOKING_UPDATED", booking)

	return &booking, nil
}

func (s *BookingService) CancelBooking(ctx context.Context, id string) error {
	var booking Booking

	err := s.db.QueryRow(ctx,
		`UPDATE bookings
		 SET status='CANCELLED',
		     updated_at=NOW()
		 WHERE id=$1
		 RETURNING id, customer_id, vehicle_id, service_package_id, booking_date, booking_slot,
		 pickup_required, pickup_address, status, created_at, updated_at`,
		id,
	).Scan(
		&booking.ID,
		&booking.CustomerID,
		&booking.VehicleID,
		&booking.ServicePackageID,
		&booking.BookingDate,
		&booking.BookingSlot,
		&booking.PickupRequired,
		&booking.PickupAddress,
		&booking.Status,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return err
	}

	_ = s.publishEvent(ctx, "BOOKING_CANCELLED", booking)

	return nil
}

func (s *BookingService) CheckSlotAvailability(ctx context.Context, date string, slot string) (bool, error) {
	var count int

	err := s.db.QueryRow(ctx,
		`SELECT COUNT(*)
		 FROM bookings
		 WHERE booking_date = $1
		   AND booking_slot = $2
		   AND status != 'CANCELLED'`,
		date,
		slot,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	maxCapacity := 5

	return count < maxCapacity, nil
}

func (s *BookingService) publishEvent(ctx context.Context, eventType string, booking Booking) error {
	event := map[string]interface{}{
		"event_type": eventType,
		"source":     "booking-service",
		"booking":    booking,
		"timestamp":  time.Now(),
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return s.redis.Publish(ctx, s.cfg.BookingEventQueue, payload).Err()
}
