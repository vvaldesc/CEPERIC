package handler

import (
	"strconv"

	"github.com/ceperic/backend/internal/domain"
	"github.com/ceperic/backend/internal/service"
	"github.com/ceperic/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, users, "Users retrieved successfully")
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid ID")
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "User not found")
	}

	return response.Success(c, user, "User retrieved successfully")
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var dto domain.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.service.CreateUser(dto)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Created(c, user, "User created successfully")
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid ID")
	}

	var dto domain.UpdateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.service.UpdateUser(uint(id), dto)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, user, "User updated successfully")
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid ID")
	}

	if err := h.service.DeleteUser(uint(id)); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, nil, "User deleted successfully")
}
