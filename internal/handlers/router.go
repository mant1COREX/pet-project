package handlers

import (
	"github.com/mant1COREX/pet-project/internal/service"
	"github.com/gofiber/fiber/v2"
)

// func RegisterProductRoutes(app *fiber.App) {
// 	api := app.Group("/api")

// 	api.Get("/products", handlers.GetProducts)          // Получить все продукты
// 	api.Post("/products", handlers.CreateProduct)       // Создать новый продукт
// 	api.Get("/products/:id", handlers.GetProduct)       // Получить продукт по ID
// 	api.Put("/products/:id", handlers.UpdateProduct)    // Обновить продукт
// 	api.Delete("/products/:id", handlers.DeleteProduct) // Удалить продукт

// }

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	tasks := app.Group("/tasks")

	tasks.Get("", h.GetAllTasks)
	tasks.Post("", h.CreateTask)
	tasks.Put("/:id", h.UpdateTask)
	tasks.Delete("/:id", h.DeleteTask)
}
