CREATE TABLE IF NOT EXISTS employee_shifts (
    id UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    shift_date VARCHAR(20) NOT NULL,
    start_time VARCHAR(20) NOT NULL,
    end_time VARCHAR(20) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'SCHEDULED',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS job_assignments (
    id UUID PRIMARY KEY,
    booking_id UUID NOT NULL,
    employee_id UUID NOT NULL,
    job_type VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'ASSIGNED',
    assigned_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS employee_attendance (
    id UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    attendance_date VARCHAR(20) NOT NULL DEFAULT CURRENT_DATE::TEXT,
    check_in_time TIMESTAMP,
    check_out_time TIMESTAMP,
    status VARCHAR(50) NOT NULL DEFAULT 'CHECKED_IN',
    UNIQUE(employee_id, attendance_date)
);

CREATE TABLE IF NOT EXISTS driver_locations (
    employee_id UUID PRIMARY KEY,
    latitude NUMERIC(10,7) NOT NULL,
    longitude NUMERIC(10,7) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_employee_shifts_employee_id
ON employee_shifts(employee_id);

CREATE INDEX IF NOT EXISTS idx_employee_shifts_date
ON employee_shifts(shift_date);

CREATE INDEX IF NOT EXISTS idx_job_assignments_employee_id
ON job_assignments(employee_id);

CREATE INDEX IF NOT EXISTS idx_job_assignments_booking_id
ON job_assignments(booking_id);

CREATE INDEX IF NOT EXISTS idx_job_assignments_status
ON job_assignments(status);

CREATE INDEX IF NOT EXISTS idx_employee_attendance_employee_date
ON employee_attendance(employee_id, attendance_date);