package internal

import "time"

type BookingEventEnvelope struct {
	EventType string           `json:"event_type"`
	Source    string           `json:"source"`
	Booking   AnalyticsBooking `json:"booking"`
	Timestamp time.Time        `json:"timestamp"`
}

type AnalyticsBooking struct {
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

type RevenueSummary struct {
	TotalBookings uint64 `json:"total_bookings"`
	TotalRevenue  uint64 `json:"total_revenue"`
}

type BookingSummary struct {
	TotalBookings     uint64 `json:"total_bookings"`
	CreatedBookings   uint64 `json:"created_bookings"`
	ConfirmedBookings uint64 `json:"confirmed_bookings"`
	CancelledBookings uint64 `json:"cancelled_bookings"`
}

type EmployeePerformanceSummary struct {
	EmployeeID     string  `json:"employee_id"`
	TotalJobs      uint64  `json:"total_jobs"`
	CompletedJobs  uint64  `json:"completed_jobs"`
	CompletionRate float64 `json:"completion_rate"`
}

type ServicePopularitySummary struct {
	ServicePackageID string `json:"service_package_id"`
	TotalBookings    uint64 `json:"total_bookings"`
}
