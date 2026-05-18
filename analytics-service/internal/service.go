package internal

import (
	"context"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
)

type AnalyticsService struct {
	ch clickhouse.Conn
}

func NewAnalyticsService(ch clickhouse.Conn) *AnalyticsService {
	return &AnalyticsService{ch: ch}
}

func (s *AnalyticsService) InsertBookingEvent(ctx context.Context, event BookingEventEnvelope) error {
	b := event.Booking

	return s.ch.Exec(ctx,
		`INSERT INTO analytics_bookings
		(event_type, source, booking_id, customer_id, vehicle_id, service_package_id,
		 booking_date, booking_slot, pickup_required, status, event_time)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		event.EventType,
		event.Source,
		b.ID,
		b.CustomerID,
		b.VehicleID,
		b.ServicePackageID,
		b.BookingDate,
		b.BookingSlot,
		b.PickupRequired,
		b.Status,
		event.Timestamp,
	)
}

func (s *AnalyticsService) GetRevenue(ctx context.Context) (*RevenueSummary, error) {
	var result RevenueSummary

	err := s.ch.QueryRow(ctx,
		`SELECT
			count() AS total_bookings,
			countIf(status != 'CANCELLED') * 1000 AS total_revenue
		 FROM analytics_bookings`,
	).Scan(&result.TotalBookings, &result.TotalRevenue)

	return &result, err
}

func (s *AnalyticsService) GetBookingSummary(ctx context.Context) (*BookingSummary, error) {
	var result BookingSummary

	err := s.ch.QueryRow(ctx,
		`SELECT
			count() AS total_bookings,
			countIf(status = 'CREATED') AS created_bookings,
			countIf(status = 'CONFIRMED') AS confirmed_bookings,
			countIf(status = 'CANCELLED') AS cancelled_bookings
		 FROM analytics_bookings`,
	).Scan(
		&result.TotalBookings,
		&result.CreatedBookings,
		&result.ConfirmedBookings,
		&result.CancelledBookings,
	)

	return &result, err
}

func (s *AnalyticsService) GetEmployeePerformance(ctx context.Context) ([]EmployeePerformanceSummary, error) {
	rows, err := s.ch.Query(ctx,
		`SELECT
			employee_id,
			count() AS total_jobs,
			countIf(status = 'COMPLETED') AS completed_jobs,
			if(count() = 0, 0, round((countIf(status = 'COMPLETED') / count()) * 100, 2)) AS completion_rate
		 FROM analytics_employee_jobs
		 GROUP BY employee_id
		 ORDER BY total_jobs DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []EmployeePerformanceSummary

	for rows.Next() {
		var item EmployeePerformanceSummary
		if err := rows.Scan(&item.EmployeeID, &item.TotalJobs, &item.CompletedJobs, &item.CompletionRate); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, rows.Err()
}

func (s *AnalyticsService) GetServicePopularity(ctx context.Context) ([]ServicePopularitySummary, error) {
	rows, err := s.ch.Query(ctx,
		`SELECT
			service_package_id,
			count() AS total_bookings
		 FROM analytics_bookings
		 GROUP BY service_package_id
		 ORDER BY total_bookings DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ServicePopularitySummary

	for rows.Next() {
		var item ServicePopularitySummary
		if err := rows.Scan(&item.ServicePackageID, &item.TotalBookings); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, rows.Err()
}
