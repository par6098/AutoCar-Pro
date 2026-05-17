package internal

import "time"

type Package struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	VehicleType string     `json:"vehicle_type"`
	BasePrice   float64    `json:"base_price"`
	Version     int        `json:"version"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type Addon struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Version     int        `json:"version"`
	IsActive    bool       `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type PricingRule struct {
	ID              string     `json:"id"`
	RuleType        string     `json:"rule_type"`
	VehicleType     string     `json:"vehicle_type"`
	PackageID       *string    `json:"package_id,omitempty"`
	AddonID         *string    `json:"addon_id,omitempty"`
	DiscountPercent float64    `json:"discount_percent"`
	OverridePrice   *float64   `json:"override_price,omitempty"`
	SeasonStart     *time.Time `json:"season_start,omitempty"`
	SeasonEnd       *time.Time `json:"season_end,omitempty"`
	IsActive        bool       `json:"is_active"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type CreatePackageRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	VehicleType string  `json:"vehicle_type"`
	BasePrice   float64 `json:"base_price"`
}

type CreateAddonRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type CreatePricingRuleRequest struct {
	RuleType        string     `json:"rule_type"`
	VehicleType     string     `json:"vehicle_type"`
	PackageID       *string    `json:"package_id"`
	AddonID         *string    `json:"addon_id"`
	DiscountPercent float64    `json:"discount_percent"`
	OverridePrice   *float64   `json:"override_price"`
	SeasonStart     *time.Time `json:"season_start"`
	SeasonEnd       *time.Time `json:"season_end"`
}

type CalculatePriceRequest struct {
	PackageID   string   `json:"package_id"`
	AddonIDs    []string `json:"addon_ids"`
	VehicleType string   `json:"vehicle_type"`
	BookingDate string   `json:"booking_date"`
}

type CalculatePriceResponse struct {
	BasePrice        float64  `json:"base_price"`
	AddonTotal       float64  `json:"addon_total"`
	DiscountAmount   float64  `json:"discount_amount"`
	SeasonalOverride *float64 `json:"seasonal_override,omitempty"`
	FinalPrice       float64  `json:"final_price"`
}
