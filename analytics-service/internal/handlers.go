package internal

import "github.com/gofiber/fiber/v2"

type AnalyticsHandler struct {
	service *AnalyticsService
}

func NewAnalyticsHandler(service *AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{service: service}
}

func (h *AnalyticsHandler) GetRevenue(c *fiber.Ctx) error {
	result, err := h.service.GetRevenue(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *AnalyticsHandler) GetBookingSummary(c *fiber.Ctx) error {
	result, err := h.service.GetBookingSummary(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *AnalyticsHandler) GetEmployeePerformance(c *fiber.Ctx) error {
	result, err := h.service.GetEmployeePerformance(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *AnalyticsHandler) GetServicePopularity(c *fiber.Ctx) error {
	result, err := h.service.GetServicePopularity(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}
