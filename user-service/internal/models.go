package internal

import "time"

type Customer struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vehicle struct {
	ID          string    `json:"id"`
	CustomerID  string    `json:"customer_id"`
	VehicleType string    `json:"vehicle_type"`
	Make        string    `json:"make"`
	Model       string    `json:"model"`
	RegNumber   string    `json:"reg_number"`
	FuelType    string    `json:"fuel_type"`
	CreatedAt   time.Time `json:"created_at"`
}

type Preference struct {
	CustomerID        string    `json:"customer_id"`
	PreferredLanguage string    `json:"preferred_language"`
	NotificationMode  string    `json:"notification_mode"`
	PickupRequired    bool      `json:"pickup_required"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Employee struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Admin struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type Tenant struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	OwnerTenantID *string   `json:"owner_tenant_id,omitempty"`
	TenantType    string    `json:"tenant_type"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateCustomerRequest struct {
	TenantID string `json:"tenant_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

type AddVehicleRequest struct {
	VehicleType string `json:"vehicle_type"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	RegNumber   string `json:"reg_number"`
	FuelType    string `json:"fuel_type"`
}

type UpdatePreferenceRequest struct {
	PreferredLanguage string `json:"preferred_language"`
	NotificationMode  string `json:"notification_mode"`
	PickupRequired    bool   `json:"pickup_required"`
}

type CreateEmployeeRequest struct {
	TenantID string `json:"tenant_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Role     string `json:"role"`
}

type AddSkillRequest struct {
	SkillTag string `json:"skill_tag"`
}

type CreateAdminRequest struct {
	TenantID string `json:"tenant_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type CreateTenantRequest struct {
	Name          string  `json:"name"`
	OwnerTenantID *string `json:"owner_tenant_id"`
	TenantType    string  `json:"tenant_type"`
}
