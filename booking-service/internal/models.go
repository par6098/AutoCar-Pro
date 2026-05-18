package internal

import "time"

type Booking struct {
	ID               string    `json:"id"`
	CustomerID       string    `json:"customer_id"`
	VehicleID        string    `json:"vehicle_id"`
	ServicePackageID string    `json:"service_package_id"`
	BookingDate      string    `json:"booking_date"`
	BookingSlot      string    `json:"booking_slot"`
	PickupRequired   bool      `json:"pickup_required"`
	PickupAddress    string    `json:"pickup_address"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreateBookingRequest struct {
	CustomerID       string `json:"customer_id"`
	VehicleID        string `json:"vehicle_id"`
	ServicePackageID string `json:"service_package_id"`
	BookingDate      string `json:"booking_date"`
	BookingSlot      string `json:"booking_slot"`
	PickupRequired   bool   `json:"pickup_required"`
	PickupAddress    string `json:"pickup_address"`
}

type UpdateBookingRequest struct {
	BookingDate    string `json:"booking_date"`
	BookingSlot    string `json:"booking_slot"`
	PickupRequired bool   `json:"pickup_required"`
	PickupAddress  string `json:"pickup_address"`
	Status         string `json:"status"`
}

const (
	BookingStatusCreated   = "CREATED"
	BookingStatusConfirmed = "CONFIRMED"
	BookingStatusCancelled = "CANCELLED"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	Status       string
	CreatedAt    time.Time
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}
