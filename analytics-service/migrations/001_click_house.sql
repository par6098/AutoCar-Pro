CREATE DATABASE IF NOT EXISTS autocare_analytics;

CREATE TABLE IF NOT EXISTS autocare_analytics.analytics_bookings (
    event_type String,
    source String,
    booking_id String,
    customer_id String,
    vehicle_id String,
    service_package_id String,
    booking_date String,
    booking_slot String,
    pickup_required Bool,
    status String,
    event_time DateTime
)
ENGINE = MergeTree()
ORDER BY (event_time, booking_id);

CREATE TABLE IF NOT EXISTS autocare_analytics.analytics_employee_jobs (
    employee_id String,
    booking_id String,
    job_type String,
    status String,
    event_time DateTime
)
ENGINE = MergeTree()
ORDER BY (event_time, employee_id);