package internal

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CatalogService struct {
	db *pgxpool.Pool
}

func NewCatalogService(db *pgxpool.Pool) *CatalogService {
	return &CatalogService{
		db: db,
	}
}

func (s *CatalogService) CreatePackage(
	ctx context.Context,
	req CreatePackageRequest,
) (*ServicePackage, error) {

	id := uuid.New().String()

	var result ServicePackage

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO service_packages (
			id,
			name,
			description,
			vehicle_type,
			base_price,
			duration_minutes,
			active,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,
			true,
			NOW()
		)
		RETURNING
			id,
			name,
			description,
			vehicle_type,
			base_price,
			duration_minutes,
			active,
			created_at
		`,
		id,
		req.Name,
		req.Description,
		req.VehicleType,
		req.BasePrice,
		req.DurationMinutes,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.VehicleType,
		&result.BasePrice,
		&result.DurationMinutes,
		&result.Active,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CatalogService) ListPackages(
	ctx context.Context,
) ([]ServicePackage, error) {

	rows, err := s.db.Query(ctx,
		`
		SELECT
			id,
			name,
			description,
			vehicle_type,
			base_price,
			duration_minutes,
			active,
			created_at
		FROM service_packages
		WHERE active = true
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []ServicePackage

	for rows.Next() {

		var item ServicePackage

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.VehicleType,
			&item.BasePrice,
			&item.DurationMinutes,
			&item.Active,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, rows.Err()
}

func (s *CatalogService) GetPackage(
	ctx context.Context,
	id string,
) (*ServicePackage, error) {

	var result ServicePackage

	err := s.db.QueryRow(ctx,
		`
		SELECT
			id,
			name,
			description,
			vehicle_type,
			base_price,
			duration_minutes,
			active,
			created_at
		FROM service_packages
		WHERE id = $1
		`,
		id,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.VehicleType,
		&result.BasePrice,
		&result.DurationMinutes,
		&result.Active,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CatalogService) UpdatePackage(
	ctx context.Context,
	id string,
	req UpdatePackageRequest,
) (*ServicePackage, error) {

	var result ServicePackage

	err := s.db.QueryRow(ctx,
		`
		UPDATE service_packages
		SET
			name = $1,
			description = $2,
			vehicle_type = $3,
			base_price = $4,
			duration_minutes = $5
		WHERE id = $6
		RETURNING
			id,
			name,
			description,
			vehicle_type,
			base_price,
			duration_minutes,
			active,
			created_at
		`,
		req.Name,
		req.Description,
		req.VehicleType,
		req.BasePrice,
		req.DurationMinutes,
		id,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.VehicleType,
		&result.BasePrice,
		&result.DurationMinutes,
		&result.Active,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CatalogService) SoftDeletePackage(
	ctx context.Context,
	id string,
) error {

	_, err := s.db.Exec(ctx,
		`
		UPDATE service_packages
		SET active = false
		WHERE id = $1
		`,
		id,
	)

	return err
}

func (s *CatalogService) CreateAddon(
	ctx context.Context,
	req CreateAddonRequest,
) (*Addon, error) {

	id := uuid.New().String()

	var result Addon

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO addons (
			id,
			name,
			price,
			active,
			created_at
		)
		VALUES (
			$1,$2,$3,
			true,
			NOW()
		)
		RETURNING
			id,
			name,
			price,
			active,
			created_at
		`,
		id,
		req.Name,
		req.Price,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Price,
		&result.Active,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CatalogService) ListAddons(
	ctx context.Context,
) ([]Addon, error) {

	rows, err := s.db.Query(ctx,
		`
		SELECT
			id,
			name,
			price,
			active,
			created_at
		FROM addons
		WHERE active = true
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []Addon

	for rows.Next() {

		var item Addon

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
			&item.Active,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, rows.Err()
}

func (s *CatalogService) SoftDeleteAddon(
	ctx context.Context,
	id string,
) error {

	_, err := s.db.Exec(ctx,
		`
		UPDATE addons
		SET active = false
		WHERE id = $1
		`,
		id,
	)

	return err
}

func (s *CatalogService) CreatePricingRule(
	ctx context.Context,
	req CreatePricingRuleRequest,
) (*PricingRule, error) {

	id := uuid.New().String()

	var result PricingRule

	err := s.db.QueryRow(ctx,
		`
		INSERT INTO pricing_rules (
			id,
			vehicle_type,
			season,
			price_multiplier,
			created_at
		)
		VALUES (
			$1,$2,$3,$4,NOW()
		)
		RETURNING
			id,
			vehicle_type,
			season,
			price_multiplier,
			created_at
		`,
		id,
		req.VehicleType,
		req.Season,
		req.PriceMultiplier,
	).Scan(
		&result.ID,
		&result.VehicleType,
		&result.Season,
		&result.PriceMultiplier,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *CatalogService) ListPricingRules(
	ctx context.Context,
) ([]PricingRule, error) {

	rows, err := s.db.Query(ctx,
		`
		SELECT
			id,
			vehicle_type,
			season,
			price_multiplier,
			created_at
		FROM pricing_rules
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []PricingRule

	for rows.Next() {

		var item PricingRule

		err := rows.Scan(
			&item.ID,
			&item.VehicleType,
			&item.Season,
			&item.PriceMultiplier,
			&item.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, rows.Err()
}

func (s *CatalogService) CalculatePrice(
	ctx context.Context,
	req CalculatePriceRequest,
) (*PriceCalculationResponse, error) {

	var packagePrice float64

	err := s.db.QueryRow(ctx,
		`
		SELECT base_price
		FROM service_packages
		WHERE id = $1
		`,
		req.PackageID,
	).Scan(&packagePrice)

	if err != nil {
		return nil, err
	}

	addonPrice := 0.0

	for _, addonID := range req.Addons {

		var price float64

		err := s.db.QueryRow(ctx,
			`
			SELECT price
			FROM addons
			WHERE id = $1
			`,
			addonID,
		).Scan(&price)

		if err != nil {
			return nil, err
		}

		addonPrice += price
	}

	multiplier := 1.0

	err = s.db.QueryRow(ctx,
		`
		SELECT price_multiplier
		FROM pricing_rules
		WHERE vehicle_type = $1
		  AND season = $2
		LIMIT 1
		`,
		req.VehicleType,
		req.Season,
	).Scan(&multiplier)

	if err != nil {
		multiplier = 1.0
	}

	total := (packagePrice + addonPrice) * multiplier

	return &PriceCalculationResponse{
		BasePrice:        packagePrice,
		AddonPrice:       addonPrice,
		SeasonMultiplier: multiplier,
		FinalPrice:       total,
	}, nil
}
