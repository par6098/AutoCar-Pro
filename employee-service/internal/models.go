package internal

import "time"

type Shift struct {
	ID         string    `json:"id"`
	EmployeeID string    `json:"employee_id"`
	ShiftDate  string    `json:"shift_date"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

type JobAssignment struct {
	ID         string    `json:"id"`
	BookingID  string    `json:"booking_id"`
	EmployeeID string    `json:"employee_id"`
	JobType    string    `json:"job_type"`
	Status     string    `json:"status"`
	AssignedAt time.Time `json:"assigned_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Attendance struct {
	ID             string     `json:"id"`
	EmployeeID     string     `json:"employee_id"`
	AttendanceDate string     `json:"attendance_date"`
	CheckInTime    *time.Time `json:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time"`
	Status         string     `json:"status"`
}

type DriverLocation struct {
	EmployeeID string    `json:"employee_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Performance struct {
	EmployeeID     string  `json:"employee_id"`
	TotalJobs      int     `json:"total_jobs"`
	CompletedJobs  int     `json:"completed_jobs"`
	CancelledJobs  int     `json:"cancelled_jobs"`
	CompletionRate float64 `json:"completion_rate"`
}

type CreateShiftRequest struct {
	EmployeeID string `json:"employee_id"`
	ShiftDate  string `json:"shift_date"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type AssignJobRequest struct {
	BookingID  string `json:"booking_id"`
	EmployeeID string `json:"employee_id"`
	JobType    string `json:"job_type"`
}

type UpdateJobStatusRequest struct {
	Status string `json:"status"`
}

type AttendanceRequest struct {
	EmployeeID string `json:"employee_id"`
}

type UpdateDriverLocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
