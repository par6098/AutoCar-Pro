package internal

import "github.com/gofiber/fiber/v2"

type BookingHandler struct {
	service *BookingService
}

func NewBookingHandler(service *BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	var req CreateBookingRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	booking, err := h.service.CreateBooking(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(booking)
}

func (h *BookingHandler) GetBooking(c *fiber.Ctx) error {
	booking, err := h.service.GetBooking(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "booking not found"})
	}

	return c.JSON(booking)
}

func (h *BookingHandler) UpdateBooking(c *fiber.Ctx) error {
	var req UpdateBookingRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	booking, err := h.service.UpdateBooking(c.Context(), c.Params("id"), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(booking)
}

func (h *BookingHandler) CancelBooking(c *fiber.Ctx) error {
	err := h.service.CancelBooking(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "booking cancelled successfully"})
}

func (h *BookingHandler) CheckSlotAvailability(c *fiber.Ctx) error {
	date := c.Query("date")
	slot := c.Query("slot")

	if date == "" || slot == "" {
		return c.Status(400).JSON(fiber.Map{"error": "date and slot are required"})
	}

	available, err := h.service.CheckSlotAvailability(c.Context(), date, slot)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"date":      date,
		"slot":      slot,
		"available": available,
	})
}
