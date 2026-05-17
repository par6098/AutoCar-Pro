package internal

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CatalogService struct {
	db *pgxpool.Pool
}

func NewCatalogService(db *pgxpool.Pool) *CatalogService {
	return &CatalogService{db: db}
}

func (s *CatalogService) CreatePackage(ctx context.Context, req CreatePackageRequest) (*Package, error) {
	id := uuid.New().String()

	var pkg Package

	err := s.db.QueryRow(ctx,
		`INSERT INTO service_packages
		(id, name, description, vehicle_type, base_price, version, is_active, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,1,true,NOW(),NOW())
		RETURNING id, name, description, vehicle_type, base_price, version, is_active, created_at, updated_at, deleted_at`,
		id, req.Name, req.Description, req.VehicleType, req.BasePrice,
	).Scan(
		&pkg.ID,
		&pkg.Name,
		&pkg.Description,
		&pkg.VehicleType,
		&pkg.BasePrice,
		&pkg.Version,
		&pkg.IsActive,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
		&pkg.DeletedAt,
	)

	return &pkg, err
}

func (s *CatalogService) ListPackages(ctx context.Context) ([]Package, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, name, description, vehicle_type, base_price, version, is_active, created_at, updated_at, deleted_at
		 FROM service_packages
		 WHERE deleted_at IS NULL AND is_active = true
		 ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []Package

	for rows.Next() {
		var pkg Package
		if err := rows.Scan(
			&pkg.ID,
			&pkg.Name,
			&pkg.Description,
			&pkg.VehicleType,
			&pkg.BasePrice,
			&pkg.Version,
			&pkg.IsActive,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
			&pkg.DeletedAt,
		); err != nil {
			return nil, err
		}

		packages = append(packages, pkg)
	}

	return packages, rows.Err()
}

func (s *CatalogService) GetPackage(ctx context.Context, id string) (*Package, error) {
	var pkg Package

	err := s.db.QueryRow(ctx,
		`SELECT id, name, description, vehicle_type, base_price, version, is_active, created_at, updated_at, deleted_at
		 FROM service_packages
		 WHERE id = $1 AND deleted_at IS NULL`,
		id,
	).Scan(
		&pkg.ID,
		&pkg.Name,
		&pkg.Description,
		&pkg.VehicleType,
		&pkg.BasePrice,
		&pkg.Version,
		&pkg.IsActive,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
		&pkg.DeletedAt,
	)

	return &pkg, err
}

func (s *CatalogService) UpdatePackage(ctx context.Context, id string, req CreatePackageRequest) (*Package, error) {
	var pkg Package

	err := s.db.QueryRow(ctx,
		`UPDATE service_packages
		 SET name=$1,
		     description=$2,
		     vehicle_type=$3,
		     base_price=$4,
		     version=version+1,
		     updated_at=NOW()
		 WHERE id=$5 AND deleted_at IS NULL
		 RETURNING id, name, description, vehicle_type, base_price, version, is_active, created_at, updated_at, deleted_at`,
		req.Name,
		req.Description,
		req.VehicleType,
		req.BasePrice,
		id,
	).Scan(
		&pkg.ID,
		&pkg.Name,
		&pkg.Description,
		&pkg.VehicleType,
		&pkg.BasePrice,
		&pkg.Version,
		&pkg.IsActive,
		&pkg.CreatedAt,
		&pkg.UpdatedAt,
		&pkg.DeletedAt,
	)

	return &pkg, err
}

func (s *CatalogService) SoftDeletePackage(ctx context.Context, id string) error {
	_, err := s.db.Exec(ctx,
		`UPDATE service_packages
		 SET deleted_at=NOW(),
		     is_active=false,
		     updated_at=NOW()
		 WHERE id=$1 AND deleted_at IS NULL`,
		id,
	)

	return err
}

func (s *CatalogService) CreateAddon(ctx context.Context, req CreateAddonRequest) (*Addon, error) {
	id := uuid.New().String()

	var addon Addon

	err := s.db.QueryRow(ctx,
		`INSERT INTO addons
		(id, name, description, price, version, is_active, created_at, updated_at)
		VALUES ($1,$2,$3,$4,1,true,NOW(),NOW())
		RETURNING id, name, description, price, version, is_active, created_at, updated_at, deleted_at`,
		id, req.Name, req.Description, req.Price,
	).Scan(
		&addon.ID,
		&addon.Name,
		&addon.Description,
		&addon.Price,
		&addon.Version,
		&addon.IsActive,
		&addon.CreatedAt,
		&addon.UpdatedAt,
		&addon.DeletedAt,
	)

	return &addon, err
}

func (s *CatalogService) ListAddons(ctx context.Context) ([]Addon, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, name, description, price, version, is_active, created_at, updated_at, deleted_at
		 FROM addons
		 WHERE deleted_at IS NULL AND is_active = true
		 ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addons []Addon

	for rows.Next() {
		var addon Addon
		if err := rows.Scan(
			&addon.ID,
			&addon.Name,
			&addon.Description,
			&addon.Price,
			&addon.Version,
			&addon.IsActive,
			&addon.CreatedAt,
			&addon.UpdatedAt,
			&addon.DeletedAt,
		); err != nil {
			return nil, err
		}

		addons = append(addons, addon)
	}

	return addons, rows.Err()
}

func (s *CatalogService) SoftDeleteAddon(ctx context.Context, id string) error {
	_, err := s.db.Exec(ctx,
		`UPDATE addons
		 SET deleted_at=NOW(),
		     is_active=false,
		     updated_at=NOW()
		 WHERE id=$1 AND deleted_at IS NULL`,
		id,
	)

	return err
}

func (s *CatalogService) CreatePricingRule(ctx context.Context, req CreatePricingRuleRequest) (*PricingRule, error) {
	id := uuid.New().String()

	var rule PricingRule

	err := s.db.QueryRow(ctx,
		`INSERT INTO pricing_rules
		(id, rule_type, vehicle_type, package_id, addon_id, discount_percent, override_price, season_start, season_end, is_active, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,true,NOW(),NOW())
		RETURNING id, rule_type, vehicle_type, package_id, addon_id, discount_percent, override_price, season_start, season_end, is_active, created_at, updated_at, deleted_at`,
		id,
		req.RuleType,
		req.VehicleType,
		req.PackageID,
		req.AddonID,
		req.DiscountPercent,
		req.OverridePrice,
		req.SeasonStart,
		req.SeasonEnd,
	).Scan(
		&rule.ID,
		&rule.RuleType,
		&rule.VehicleType,
		&rule.PackageID,
		&rule.AddonID,
		&rule.DiscountPercent,
		&rule.OverridePrice,
		&rule.SeasonStart,
		&rule.SeasonEnd,
		&rule.IsActive,
		&rule.CreatedAt,
		&rule.UpdatedAt,
		&rule.DeletedAt,
	)

	return &rule, err
}

func (s *CatalogService) ListPricingRules(ctx context.Context) ([]PricingRule, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, rule_type, vehicle_type, package_id, addon_id, discount_percent, override_price, season_start, season_end, is_active, created_at, updated_at, deleted_at
		 FROM pricing_rules
		 WHERE deleted_at IS NULL AND is_active = true
		 ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []PricingRule

	for rows.Next() {
		var rule PricingRule
		if err := rows.Scan(
			&rule.ID,
			&rule.RuleType,
			&rule.VehicleType,
			&rule.PackageID,
			&rule.AddonID,
			&rule.DiscountPercent,
			&rule.OverridePrice,
			&rule.SeasonStart,
			&rule.SeasonEnd,
			&rule.IsActive,
			&rule.CreatedAt,
			&rule.UpdatedAt,
			&rule.DeletedAt,
		); err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	return rules, rows.Err()
}

func (s *CatalogService) CalculatePrice(ctx context.Context, req CalculatePriceRequest) (*CalculatePriceResponse, error) {
	bookingDate, err := time.Parse(time.RFC3339, req.BookingDate)
	if err != nil {
		return nil, errors.New("invalid booking_date, expected RFC3339 format")
	}

	var basePrice float64

	err = s.db.QueryRow(ctx,
		`SELECT base_price
		 FROM service_packages
		 WHERE id=$1
		   AND vehicle_type=$2
		   AND deleted_at IS NULL
		   AND is_active=true`,
		req.PackageID,
		req.VehicleType,
	).Scan(&basePrice)

	if err != nil {
		return nil, errors.New("package not found")
	}

	addonTotal, err := s.getAddonTotal(ctx, req.AddonIDs)
	if err != nil {
		return nil, err
	}

	total := basePrice + addonTotal

	seasonalOverride, err := s.getSeasonalOverride(ctx, req.PackageID, req.VehicleType, bookingDate)
	if err != nil {
		return nil, err
	}

	if seasonalOverride != nil {
		total = *seasonalOverride + addonTotal
	}

	discountPercent, err := s.getBestDiscount(ctx, req.PackageID, req.VehicleType, bookingDate)
	if err != nil {
		return nil, err
	}

	discountAmount := total * discountPercent / 100
	finalPrice := total - discountAmount

	return &CalculatePriceResponse{
		BasePrice:        basePrice,
		AddonTotal:       addonTotal,
		DiscountAmount:   discountAmount,
		SeasonalOverride: seasonalOverride,
		FinalPrice:       finalPrice,
	}, nil
}

func (s *CatalogService) getAddonTotal(ctx context.Context, addonIDs []string) (float64, error) {
	if len(addonIDs) == 0 {
		return 0, nil
	}

	var total float64

	err := s.db.QueryRow(ctx,
		`SELECT COALESCE(SUM(price), 0)
		 FROM addons
		 WHERE id = ANY($1)
		   AND deleted_at IS NULL
		   AND is_active = true`,
		addonIDs,
	).Scan(&total)

	return total, err
}

func (s *CatalogService) getSeasonalOverride(ctx context.Context, packageID, vehicleType string, bookingDate time.Time) (*float64, error) {
	var overridePrice float64

	err := s.db.QueryRow(ctx,
		`SELECT override_price
		 FROM pricing_rules
		 WHERE rule_type = 'SEASONAL_OVERRIDE'
		   AND package_id = $1
		   AND vehicle_type = $2
		   AND override_price IS NOT NULL
		   AND season_start <= $3
		   AND season_end >= $3
		   AND deleted_at IS NULL
		   AND is_active = true
		 ORDER BY created_at DESC
		 LIMIT 1`,
		packageID,
		vehicleType,
		bookingDate,
	).Scan(&overridePrice)

	if err != nil {
		return nil, nil
	}

	return &overridePrice, nil
}

func (s *CatalogService) getBestDiscount(ctx context.Context, packageID, vehicleType string, bookingDate time.Time) (float64, error) {
	var discount float64

	err := s.db.QueryRow(ctx,
		`SELECT COALESCE(MAX(discount_percent), 0)
		 FROM pricing_rules
		 WHERE rule_type IN ('BUNDLE_DISCOUNT', 'SEASONAL_DISCOUNT')
		   AND package_id = $1
		   AND vehicle_type = $2
		   AND deleted_at IS NULL
		   AND is_active = true
		   AND (
		        season_start IS NULL
		        OR season_end IS NULL
		        OR (season_start <= $3 AND season_end >= $3)
		   )`,
		packageID,
		vehicleType,
		bookingDate,
	).Scan(&discount)

	return discount, err
}
