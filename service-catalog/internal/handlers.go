package internal

import "github.com/gofiber/fiber/v3"

type CatalogHandler struct {
	service *CatalogService
}

func NewCatalogHandler(service *CatalogService) *CatalogHandler {
	return &CatalogHandler{service: service}
}

func (h *CatalogHandler) CreatePackage(c fiber.Ctx) error {
	var req CreatePackageRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	pkg, err := h.service.CreatePackage(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(pkg)
}

func (h *CatalogHandler) ListPackages(c fiber.Ctx) error {
	items, err := h.service.ListPackages(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *CatalogHandler) GetPackage(c fiber.Ctx) error {
	item, err := h.service.GetPackage(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "package not found"})
	}
	return c.JSON(item)
}

func (h *CatalogHandler) UpdatePackage(c fiber.Ctx) error {
	var req CreatePackageRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	item, err := h.service.UpdatePackage(c.Context(), c.Params("id"), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(item)
}

func (h *CatalogHandler) SoftDeletePackage(c fiber.Ctx) error {
	if err := h.service.SoftDeletePackage(c.Context(), c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "package soft deleted"})
}

func (h *CatalogHandler) CreateAddon(c fiber.Ctx) error {
	var req CreateAddonRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	item, err := h.service.CreateAddon(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(item)
}

func (h *CatalogHandler) ListAddons(c fiber.Ctx) error {
	items, err := h.service.ListAddons(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *CatalogHandler) SoftDeleteAddon(c fiber.Ctx) error {
	if err := h.service.SoftDeleteAddon(c.Context(), c.Params("id")); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "addon soft deleted"})
}

func (h *CatalogHandler) CreatePricingRule(c fiber.Ctx) error {
	var req CreatePricingRuleRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	rule, err := h.service.CreatePricingRule(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(rule)
}

func (h *CatalogHandler) ListPricingRules(c fiber.Ctx) error {
	items, err := h.service.ListPricingRules(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *CatalogHandler) CalculatePrice(c fiber.Ctx) error {
	var req CalculatePriceRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	price, err := h.service.CalculatePrice(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(price)
}
