package internal

import "time"

type BookingStatus string

const (
	StatusCreated   BookingStatus = "CREATED"
	StatusUpdated   BookingStatus = "UPDATED"
	StatusCancelled BookingStatus = "CANCELLED"
)

type Booking struct {
	ID            string        `json:"id"`
	CustomerID    string        `json:"customer_id"`
	ServiceID     string        `json:"service_id"`
	SlotStart     time.Time     `json:"slot_start"`
	SlotEnd       time.Time     `json:"slot_end"`
	PickupAddress string        `json:"pickup_address"`
	DropAddress   string        `json:"drop_address"`
	PickupTime    time.Time     `json:"pickup_time"`
	DropTime      time.Time     `json:"drop_time"`
	Status        BookingStatus `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type CreateBookingRequest struct {
	CustomerID    string    `json:"customer_id"`
	ServiceID     string    `json:"service_id"`
	SlotStart     time.Time `json:"slot_start"`
	SlotEnd       time.Time `json:"slot_end"`
	PickupAddress string    `json:"pickup_address"`
	DropAddress   string    `json:"drop_address"`
	PickupTime    time.Time `json:"pickup_time"`
	DropTime      time.Time `json:"drop_time"`
}

type UpdateBookingRequest struct {
	SlotStart     time.Time `json:"slot_start"`
	SlotEnd       time.Time `json:"slot_end"`
	PickupAddress string    `json:"pickup_address"`
	DropAddress   string    `json:"drop_address"`
	PickupTime    time.Time `json:"pickup_time"`
	DropTime      time.Time `json:"drop_time"`
}

type AvailabilityResponse struct {
	Available bool   `json:"available"`
	Message   string `json:"message"`
}
