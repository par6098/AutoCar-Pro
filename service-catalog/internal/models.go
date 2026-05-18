package internal

import "time"

type ServicePackage struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	VehicleType     string    `json:"vehicle_type"`
	BasePrice       float64   `json:"base_price"`
	DurationMinutes int       `json:"duration_minutes"`
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
}

type Addon struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

type PricingRule struct {
	ID              string    `json:"id"`
	VehicleType     string    `json:"vehicle_type"`
	Season          string    `json:"season"`
	PriceMultiplier float64   `json:"price_multiplier"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreatePackageRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	VehicleType     string  `json:"vehicle_type"`
	BasePrice       float64 `json:"base_price"`
	DurationMinutes int     `json:"duration_minutes"`
}

type UpdatePackageRequest struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	VehicleType     string  `json:"vehicle_type"`
	BasePrice       float64 `json:"base_price"`
	DurationMinutes int     `json:"duration_minutes"`
}

type CreateAddonRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreatePricingRuleRequest struct {
	VehicleType     string  `json:"vehicle_type"`
	Season          string  `json:"season"`
	PriceMultiplier float64 `json:"price_multiplier"`
}

type CalculatePriceRequest struct {
	PackageID   string   `json:"package_id"`
	VehicleType string   `json:"vehicle_type"`
	Addons      []string `json:"addons"`
	Season      string   `json:"season"`
}

type PriceCalculationResponse struct {
	BasePrice        float64 `json:"base_price"`
	AddonPrice       float64 `json:"addon_price"`
	SeasonMultiplier float64 `json:"season_multiplier"`
	FinalPrice       float64 `json:"final_price"`
}
