package handlers

import (
	"context"
	"errors"
	"strconv"

	"github.com/mant1COREX/pet-project/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	var task entity.Task
	if err := c.BodyParser(&task); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	err := isTaskIputValid(task)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	id, err := h.services.CreateTask(context.Background(), task)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handler) DeleteTask(c *fiber.Ctx) error {
	strId := c.Params("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}
	if intId < 1 {
		c.Status(fiber.StatusBadRequest)
		return errors.New("invalid path value")
	}

	id, err := h.services.DeleteTask(context.Background(), intId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func (h *Handler) UpdateTask(c *fiber.Ctx) error {
	var task entity.Task
	if err := c.BodyParser(&task); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	strId := c.Params("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}
	if intId < 1 {
		c.Status(fiber.StatusBadRequest)
		return errors.New("invalid path value")
	}

	err = isTaskIputValid(task)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	task.Id = intId
	updatedTask, err := h.services.UpdateTask(context.Background(), task)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":          updatedTask.Id,
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
		"status":      updatedTask.Status,
		"created_at":  updatedTask.CreatedAt,
		"updatet_at":  updatedTask.UpdatedAt,
	})
}

func (h *Handler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.services.GetAllTasks(context.Background())
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	return c.Status(fiber.StatusOK).JSON(tasks)
}
