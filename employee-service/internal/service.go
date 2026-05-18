package internal

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeService struct {
	db *pgxpool.Pool
}

func NewEmployeeService(db *pgxpool.Pool) *EmployeeService {
	return &EmployeeService{db: db}
}

func (s *EmployeeService) CreateShift(ctx context.Context, req CreateShiftRequest) (*Shift, error) {
	id := uuid.New().String()
	var result Shift

	err := s.db.QueryRow(ctx,
		`INSERT INTO employee_shifts
		(id, employee_id, shift_date, start_time, end_time, status, created_at)
		VALUES ($1,$2,$3,$4,$5,'SCHEDULED',NOW())
		RETURNING id, employee_id, shift_date, start_time, end_time, status, created_at`,
		id, req.EmployeeID, req.ShiftDate, req.StartTime, req.EndTime,
	).Scan(&result.ID, &result.EmployeeID, &result.ShiftDate, &result.StartTime, &result.EndTime, &result.Status, &result.CreatedAt)

	return &result, err
}

func (s *EmployeeService) ListShifts(ctx context.Context, employeeID string) ([]Shift, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, employee_id, shift_date, start_time, end_time, status, created_at
		 FROM employee_shifts
		 WHERE employee_id=$1
		 ORDER BY shift_date DESC`,
		employeeID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Shift
	for rows.Next() {
		var item Shift
		if err := rows.Scan(&item.ID, &item.EmployeeID, &item.ShiftDate, &item.StartTime, &item.EndTime, &item.Status, &item.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, rows.Err()
}

func (s *EmployeeService) AssignJob(ctx context.Context, req AssignJobRequest) (*JobAssignment, error) {
	id := uuid.New().String()
	var result JobAssignment

	err := s.db.QueryRow(ctx,
		`INSERT INTO job_assignments
		(id, booking_id, employee_id, job_type, status, assigned_at, updated_at)
		VALUES ($1,$2,$3,$4,'ASSIGNED',NOW(),NOW())
		RETURNING id, booking_id, employee_id, job_type, status, assigned_at, updated_at`,
		id, req.BookingID, req.EmployeeID, req.JobType,
	).Scan(&result.ID, &result.BookingID, &result.EmployeeID, &result.JobType, &result.Status, &result.AssignedAt, &result.UpdatedAt)

	return &result, err
}

func (s *EmployeeService) UpdateJobStatus(ctx context.Context, id string, status string) (*JobAssignment, error) {
	var result JobAssignment

	err := s.db.QueryRow(ctx,
		`UPDATE job_assignments
		 SET status=$1, updated_at=NOW()
		 WHERE id=$2
		 RETURNING id, booking_id, employee_id, job_type, status, assigned_at, updated_at`,
		status, id,
	).Scan(&result.ID, &result.BookingID, &result.EmployeeID, &result.JobType, &result.Status, &result.AssignedAt, &result.UpdatedAt)

	return &result, err
}

func (s *EmployeeService) ListJobs(ctx context.Context, employeeID string) ([]JobAssignment, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, booking_id, employee_id, job_type, status, assigned_at, updated_at
		 FROM job_assignments
		 WHERE employee_id=$1
		 ORDER BY assigned_at DESC`,
		employeeID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []JobAssignment
	for rows.Next() {
		var item JobAssignment
		if err := rows.Scan(&item.ID, &item.BookingID, &item.EmployeeID, &item.JobType, &item.Status, &item.AssignedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	return result, rows.Err()
}

func (s *EmployeeService) CheckIn(ctx context.Context, employeeID string) (*Attendance, error) {
	id := uuid.New().String()
	var result Attendance

	err := s.db.QueryRow(ctx,
		`INSERT INTO employee_attendance
		(id, employee_id, attendance_date, check_in_time, status)
		VALUES ($1,$2,CURRENT_DATE,NOW(),'CHECKED_IN')
		ON CONFLICT (employee_id, attendance_date)
		DO UPDATE SET check_in_time=COALESCE(employee_attendance.check_in_time, NOW()), status='CHECKED_IN'
		RETURNING id, employee_id, attendance_date, check_in_time, check_out_time, status`,
		id, employeeID,
	).Scan(&result.ID, &result.EmployeeID, &result.AttendanceDate, &result.CheckInTime, &result.CheckOutTime, &result.Status)

	return &result, err
}

func (s *EmployeeService) CheckOut(ctx context.Context, employeeID string) (*Attendance, error) {
	var result Attendance

	err := s.db.QueryRow(ctx,
		`UPDATE employee_attendance
		 SET check_out_time=NOW(), status='CHECKED_OUT'
		 WHERE employee_id=$1 AND attendance_date=CURRENT_DATE
		 RETURNING id, employee_id, attendance_date, check_in_time, check_out_time, status`,
		employeeID,
	).Scan(&result.ID, &result.EmployeeID, &result.AttendanceDate, &result.CheckInTime, &result.CheckOutTime, &result.Status)

	return &result, err
}

func (s *EmployeeService) GetPerformance(ctx context.Context, employeeID string) (*Performance, error) {
	var result Performance

	err := s.db.QueryRow(ctx,
		`SELECT
			$1::text AS employee_id,
			COUNT(*)::int AS total_jobs,
			COUNT(*) FILTER (WHERE status='COMPLETED')::int AS completed_jobs,
			COUNT(*) FILTER (WHERE status='CANCELLED')::int AS cancelled_jobs,
			CASE
				WHEN COUNT(*) = 0 THEN 0
				ELSE ROUND((COUNT(*) FILTER (WHERE status='COMPLETED')::numeric / COUNT(*)::numeric) * 100, 2)
			END AS completion_rate
		 FROM job_assignments
		 WHERE employee_id=$1`,
		employeeID,
	).Scan(&result.EmployeeID, &result.TotalJobs, &result.CompletedJobs, &result.CancelledJobs, &result.CompletionRate)

	return &result, err
}

func (s *EmployeeService) UpdateDriverLocation(ctx context.Context, employeeID string, req UpdateDriverLocationRequest) (*DriverLocation, error) {
	var result DriverLocation

	err := s.db.QueryRow(ctx,
		`INSERT INTO driver_locations
		(employee_id, latitude, longitude, updated_at)
		VALUES ($1,$2,$3,NOW())
		ON CONFLICT (employee_id)
		DO UPDATE SET latitude=$2, longitude=$3, updated_at=NOW()
		RETURNING employee_id, latitude, longitude, updated_at`,
		employeeID, req.Latitude, req.Longitude,
	).Scan(&result.EmployeeID, &result.Latitude, &result.Longitude, &result.UpdatedAt)

	return &result, err
}

func (s *EmployeeService) GetDriverLocation(ctx context.Context, employeeID string) (*DriverLocation, error) {
	var result DriverLocation

	err := s.db.QueryRow(ctx,
		`SELECT employee_id, latitude, longitude, updated_at
		 FROM driver_locations
		 WHERE employee_id=$1`,
		employeeID,
	).Scan(&result.EmployeeID, &result.Latitude, &result.Longitude, &result.UpdatedAt)

	return &result, err
}
