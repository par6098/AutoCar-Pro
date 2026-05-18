package internal

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateTenant(c *fiber.Ctx) error {

	var req CreateTenantRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	tenant, err := h.service.CreateTenant(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(tenant)
}

func (h *UserHandler) GetTenantHierarchy(c *fiber.Ctx) error {

	id := c.Params("id")

	hierarchy, err := h.service.GetTenantHierarchy(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(hierarchy)
}

func (h *UserHandler) CreateCustomer(c *fiber.Ctx) error {

	var req CreateCustomerRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	customer, err := h.service.CreateCustomer(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(customer)
}

func (h *UserHandler) GetCustomer(c *fiber.Ctx) error {

	id := c.Params("id")

	customer, err := h.service.GetCustomer(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "customer not found",
		})
	}

	return c.JSON(customer)
}

func (h *UserHandler) AddVehicle(c *fiber.Ctx) error {

	customerID := c.Params("id")

	var req AddVehicleRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	vehicle, err := h.service.AddVehicle(
		c.Context(),
		customerID,
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(vehicle)
}

func (h *UserHandler) ListVehicles(c *fiber.Ctx) error {

	customerID := c.Params("id")

	vehicles, err := h.service.ListVehicles(
		c.Context(),
		customerID,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(vehicles)
}

func (h *UserHandler) UpdatePreferences(c *fiber.Ctx) error {

	customerID := c.Params("id")

	var req UpdatePreferenceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	preferences, err := h.service.UpdatePreferences(
		c.Context(),
		customerID,
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(preferences)
}

func (h *UserHandler) CreateEmployee(c *fiber.Ctx) error {

	var req CreateEmployeeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	employee, err := h.service.CreateEmployee(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(employee)
}

func (h *UserHandler) GetEmployee(c *fiber.Ctx) error {

	id := c.Params("id")

	employee, err := h.service.GetEmployee(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "employee not found",
		})
	}

	return c.JSON(employee)
}

func (h *UserHandler) AddEmployeeSkill(c *fiber.Ctx) error {

	employeeID := c.Params("id")

	var req AddSkillRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.service.AddEmployeeSkill(
		c.Context(),
		employeeID,
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "skill added successfully",
	})
}

func (h *UserHandler) CreateAdmin(c *fiber.Ctx) error {

	var req CreateAdminRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	admin, err := h.service.CreateAdmin(
		c.Context(),
		req,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(admin)
}
