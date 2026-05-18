package internal

import "github.com/gofiber/fiber/v2"

type EmployeeHandler struct {
	service *EmployeeService
}

func NewEmployeeHandler(service *EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateShift(c *fiber.Ctx) error {
	var req CreateShiftRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.CreateShift(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(result)
}

func (h *EmployeeHandler) ListShifts(c *fiber.Ctx) error {
	result, err := h.service.ListShifts(c.Context(), c.Params("employee_id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) AssignJob(c *fiber.Ctx) error {
	var req AssignJobRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.AssignJob(c.Context(), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(result)
}

func (h *EmployeeHandler) UpdateJobStatus(c *fiber.Ctx) error {
	var req UpdateJobStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.UpdateJobStatus(c.Context(), c.Params("id"), req.Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) ListJobs(c *fiber.Ctx) error {
	result, err := h.service.ListJobs(c.Context(), c.Params("employee_id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) CheckIn(c *fiber.Ctx) error {
	var req AttendanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.CheckIn(c.Context(), req.EmployeeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) CheckOut(c *fiber.Ctx) error {
	var req AttendanceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.CheckOut(c.Context(), req.EmployeeID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) GetPerformance(c *fiber.Ctx) error {
	result, err := h.service.GetPerformance(c.Context(), c.Params("employee_id"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) UpdateDriverLocation(c *fiber.Ctx) error {
	var req UpdateDriverLocationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.service.UpdateDriverLocation(c.Context(), c.Params("employee_id"), req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *EmployeeHandler) GetDriverLocation(c *fiber.Ctx) error {
	result, err := h.service.GetDriverLocation(c.Context(), c.Params("employee_id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "location not found"})
	}

	return c.JSON(result)
}
