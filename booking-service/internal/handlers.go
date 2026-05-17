package internal

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type BookingHandler struct {
	service *BookingService
}

func NewBookingHandler(service *BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) CreateBooking(c fiber.Ctx) error {
	var req CreateBookingRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	booking, err := h.service.CreateBooking(c.Context(), req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(booking)
}

func (h *BookingHandler) GetBooking(c fiber.Ctx) error {
	id := c.Params("id")

	booking, err := h.service.GetBooking(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "booking not found"})
	}

	return c.JSON(booking)
}

func (h *BookingHandler) UpdateBooking(c fiber.Ctx) error {
	id := c.Params("id")

	var req UpdateBookingRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	booking, err := h.service.UpdateBooking(c.Context(), id, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(booking)
}

func (h *BookingHandler) CancelBooking(c fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.CancelBooking(c.Context(), id); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "booking cancelled successfully",
	})
}

func (h *BookingHandler) CheckSlotAvailability(c fiber.Ctx) error {
	serviceID := c.Query("service_id")
	start := c.Query("slot_start")
	end := c.Query("slot_end")

	slotStart, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid slot_start"})
	}

	slotEnd, err := time.Parse(time.RFC3339, end)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid slot_end"})
	}

	available, err := h.service.IsSlotAvailable(c.Context(), serviceID, slotStart, slotEnd)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "availability check failed"})
	}

	message := "slot not available"
	if available {
		message = "slot available"
	}

	return c.JSON(AvailabilityResponse{
		Available: available,
		Message:   message,
	})
}
