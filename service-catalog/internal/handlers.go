package internal

import (
	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	service *CatalogService
}

func NewCatalogHandler(service *CatalogService) *CatalogHandler {
	return &CatalogHandler{
		service: service,
	}
}

func (h *CatalogHandler) CreatePackage(c *fiber.Ctx) error {

	var req CreatePackageRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.service.CreatePackage(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(result)
}

func (h *CatalogHandler) ListPackages(c *fiber.Ctx) error {

	result, err := h.service.ListPackages(
		c.Context(),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func (h *CatalogHandler) GetPackage(c *fiber.Ctx) error {

	id := c.Params("id")

	result, err := h.service.GetPackage(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "package not found",
		})
	}

	return c.JSON(result)
}

func (h *CatalogHandler) UpdatePackage(c *fiber.Ctx) error {

	id := c.Params("id")

	var req UpdatePackageRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.service.UpdatePackage(
		c.Context(),
		id,
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func (h *CatalogHandler) SoftDeletePackage(c *fiber.Ctx) error {

	id := c.Params("id")

	err := h.service.SoftDeletePackage(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "package deleted successfully",
	})
}

func (h *CatalogHandler) CreateAddon(c *fiber.Ctx) error {

	var req CreateAddonRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.service.CreateAddon(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(result)
}

func (h *CatalogHandler) ListAddons(c *fiber.Ctx) error {

	result, err := h.service.ListAddons(
		c.Context(),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func (h *CatalogHandler) SoftDeleteAddon(c *fiber.Ctx) error {

	id := c.Params("id")

	err := h.service.SoftDeleteAddon(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "addon deleted successfully",
	})
}

func (h *CatalogHandler) CreatePricingRule(c *fiber.Ctx) error {

	var req CreatePricingRuleRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.service.CreatePricingRule(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(result)
}

func (h *CatalogHandler) ListPricingRules(c *fiber.Ctx) error {

	result, err := h.service.ListPricingRules(
		c.Context(),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func (h *CatalogHandler) CalculatePrice(c *fiber.Ctx) error {

	var req CalculatePriceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.service.CalculatePrice(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}
