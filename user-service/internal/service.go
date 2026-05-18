package internal

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateTenant(
	ctx context.Context,
	req CreateTenantRequest,
) (*Tenant, error) {

	id := uuid.New().String()

	var tenant Tenant

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO tenants (
			id,
			name,
			owner_tenant_id,
			tenant_type,
			status,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,
			'ACTIVE',
			NOW()
		)
		RETURNING
			id,
			name,
			owner_tenant_id,
			tenant_type,
			status,
			created_at
		`,
		id,
		req.Name,
		req.OwnerTenantID,
		req.TenantType,
	).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.OwnerTenantID,
		&tenant.TenantType,
		&tenant.Status,
		&tenant.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &tenant, nil
}

func (s *UserService) GetTenantHierarchy(
	ctx context.Context,
	tenantID string,
) ([]Tenant, error) {

	rows, err := s.db.Query(ctx,
		`
		WITH RECURSIVE tenant_tree AS (

			SELECT
				id,
				name,
				owner_tenant_id,
				tenant_type,
				status,
				created_at
			FROM tenants
			WHERE id = $1

			UNION ALL

			SELECT
				t.id,
				t.name,
				t.owner_tenant_id,
				t.tenant_type,
				t.status,
				t.created_at
			FROM tenants t
			INNER JOIN tenant_tree tt
				ON t.owner_tenant_id = tt.id
		)

		SELECT
			id,
			name,
			owner_tenant_id,
			tenant_type,
			status,
			created_at
		FROM tenant_tree
		ORDER BY created_at ASC
		`,
		tenantID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tenants []Tenant

	for rows.Next() {

		var tenant Tenant

		err := rows.Scan(
			&tenant.ID,
			&tenant.Name,
			&tenant.OwnerTenantID,
			&tenant.TenantType,
			&tenant.Status,
			&tenant.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		tenants = append(tenants, tenant)
	}

	return tenants, rows.Err()
}

func (s *UserService) CreateCustomer(
	ctx context.Context,
	req CreateCustomerRequest,
) (*Customer, error) {

	id := uuid.New().String()

	var customer Customer

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO customers (
			id,
			tenant_id,
			full_name,
			email,
			mobile,
			status,
			created_at,
			updated_at
		)
		VALUES (
			$1,$2,$3,$4,$5,
			'ACTIVE',
			NOW(),
			NOW()
		)
		RETURNING
			id,
			tenant_id,
			full_name,
			email,
			mobile,
			status,
			created_at,
			updated_at
		`,
		id,
		req.TenantID,
		req.FullName,
		req.Email,
		req.Mobile,
	).Scan(
		&customer.ID,
		&customer.TenantID,
		&customer.FullName,
		&customer.Email,
		&customer.Mobile,
		&customer.Status,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *UserService) GetCustomer(
	ctx context.Context,
	id string,
) (*Customer, error) {

	var customer Customer

	err := s.db.QueryRow(ctx,
		`
		SELECT
			id,
			tenant_id,
			full_name,
			email,
			mobile,
			status,
			created_at,
			updated_at
		FROM customers
		WHERE id = $1
		`,
		id,
	).Scan(
		&customer.ID,
		&customer.TenantID,
		&customer.FullName,
		&customer.Email,
		&customer.Mobile,
		&customer.Status,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *UserService) AddVehicle(
	ctx context.Context,
	customerID string,
	req AddVehicleRequest,
) (*Vehicle, error) {

	id := uuid.New().String()

	var vehicle Vehicle

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO customer_vehicles (
			id,
			customer_id,
			vehicle_type,
			make,
			model,
			reg_number,
			fuel_type,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,NOW()
		)
		RETURNING
			id,
			customer_id,
			vehicle_type,
			make,
			model,
			reg_number,
			fuel_type,
			created_at
		`,
		id,
		customerID,
		req.VehicleType,
		req.Make,
		req.Model,
		req.RegNumber,
		req.FuelType,
	).Scan(
		&vehicle.ID,
		&vehicle.CustomerID,
		&vehicle.VehicleType,
		&vehicle.Make,
		&vehicle.Model,
		&vehicle.RegNumber,
		&vehicle.FuelType,
		&vehicle.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (s *UserService) ListVehicles(
	ctx context.Context,
	customerID string,
) ([]Vehicle, error) {

	rows, err := s.db.Query(ctx,
		`
		SELECT
			id,
			customer_id,
			vehicle_type,
			make,
			model,
			reg_number,
			fuel_type,
			created_at
		FROM customer_vehicles
		WHERE customer_id = $1
		ORDER BY created_at DESC
		`,
		customerID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var vehicles []Vehicle

	for rows.Next() {

		var vehicle Vehicle

		err := rows.Scan(
			&vehicle.ID,
			&vehicle.CustomerID,
			&vehicle.VehicleType,
			&vehicle.Make,
			&vehicle.Model,
			&vehicle.RegNumber,
			&vehicle.FuelType,
			&vehicle.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, rows.Err()
}

func (s *UserService) UpdatePreferences(
	ctx context.Context,
	customerID string,
	req UpdatePreferenceRequest,
) (*Preference, error) {

	var pref Preference

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO customer_preferences (
			customer_id,
			preferred_language,
			notification_mode,
			pickup_required,
			updated_at
		)
		VALUES (
			$1,$2,$3,$4,NOW()
		)
		ON CONFLICT (customer_id)
		DO UPDATE SET
			preferred_language = $2,
			notification_mode = $3,
			pickup_required = $4,
			updated_at = NOW()
		RETURNING
			customer_id,
			preferred_language,
			notification_mode,
			pickup_required,
			updated_at
		`,
		customerID,
		req.PreferredLanguage,
		req.NotificationMode,
		req.PickupRequired,
	).Scan(
		&pref.CustomerID,
		&pref.PreferredLanguage,
		&pref.NotificationMode,
		&pref.PickupRequired,
		&pref.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &pref, nil
}

func (s *UserService) CreateEmployee(
	ctx context.Context,
	req CreateEmployeeRequest,
) (*Employee, error) {

	id := uuid.New().String()

	var emp Employee

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO employees (
			id,
			tenant_id,
			full_name,
			email,
			mobile,
			role,
			status,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,
			'ACTIVE',
			NOW()
		)
		RETURNING
			id,
			tenant_id,
			full_name,
			email,
			mobile,
			role,
			status,
			created_at
		`,
		id,
		req.TenantID,
		req.FullName,
		req.Email,
		req.Mobile,
		req.Role,
	).Scan(
		&emp.ID,
		&emp.TenantID,
		&emp.FullName,
		&emp.Email,
		&emp.Mobile,
		&emp.Role,
		&emp.Status,
		&emp.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (s *UserService) GetEmployee(
	ctx context.Context,
	id string,
) (*Employee, error) {

	var emp Employee

	err := s.db.QueryRow(ctx,
		`
		SELECT
			id,
			tenant_id,
			full_name,
			email,
			mobile,
			role,
			status,
			created_at
		FROM employees
		WHERE id = $1
		`,
		id,
	).Scan(
		&emp.ID,
		&emp.TenantID,
		&emp.FullName,
		&emp.Email,
		&emp.Mobile,
		&emp.Role,
		&emp.Status,
		&emp.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (s *UserService) AddEmployeeSkill(
	ctx context.Context,
	employeeID string,
	req AddSkillRequest,
) error {

	_, err := s.db.Exec(ctx,
		`
		INSERT INTO employee_skills (
			id,
			employee_id,
			skill_tag,
			created_at
		)
		VALUES (
			$1,$2,$3,NOW()
		)
		`,
		uuid.New().String(),
		employeeID,
		req.SkillTag,
	)

	return err
}

func (s *UserService) CreateAdmin(
	ctx context.Context,
	req CreateAdminRequest,
) (*Admin, error) {

	id := uuid.New().String()

	var admin Admin

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO admins (
			id,
			tenant_id,
			full_name,
			email,
			role,
			status,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,$5,
			'ACTIVE',
			NOW()
		)
		RETURNING
			id,
			tenant_id,
			full_name,
			email,
			role,
			status,
			created_at
		`,
		id,
		req.TenantID,
		req.FullName,
		req.Email,
		req.Role,
	).Scan(
		&admin.ID,
		&admin.TenantID,
		&admin.FullName,
		&admin.Email,
		&admin.Role,
		&admin.Status,
		&admin.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
