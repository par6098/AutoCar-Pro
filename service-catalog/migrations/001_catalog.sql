CREATE TABLE service_packages (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    vehicle_type VARCHAR(50) NOT NULL,
    base_price NUMERIC(12,2) NOT NULL,
    version INT NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE addons (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(12,2) NOT NULL,
    version INT NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE pricing_rules (
    id UUID PRIMARY KEY,
    rule_type VARCHAR(50) NOT NULL,
    vehicle_type VARCHAR(50),
    package_id UUID,
    addon_id UUID,
    discount_percent NUMERIC(5,2) DEFAULT 0,
    override_price NUMERIC(12,2),
    season_start TIMESTAMP,
    season_end TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE INDEX idx_packages_vehicle_type ON service_packages(vehicle_type);
CREATE INDEX idx_packages_active ON service_packages(is_active, deleted_at);
CREATE INDEX idx_addons_active ON addons(is_active, deleted_at);
CREATE INDEX idx_pricing_rules_type ON pricing_rules(rule_type);
CREATE INDEX idx_pricing_rules_vehicle_type ON pricing_rules(vehicle_type);
CREATE INDEX idx_pricing_rules_active ON pricing_rules(is_active, deleted_at);