package router

import (
	"github.com/ceperic/backend/internal/handler"
	"github.com/ceperic/backend/internal/repository"
	"github.com/ceperic/backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB, firebaseApp interface{}) {
	// API v1 group
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "CEPERIC API is running",
			"version": "1.0.0",
		})
	})

	// Users routes
	setupUserRoutes(api, db)

	// Documents routes (TODO)
	// setupDocumentRoutes(api, db)
}

func setupUserRoutes(api fiber.Router, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	users := api.Group("/users")
	users.Get("/", userHandler.GetAll)
	users.Get("/:id", userHandler.GetByID)
	users.Post("/", userHandler.Create)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)
}
