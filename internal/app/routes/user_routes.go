package routes

import (
	"github.com/ESD-Laboratory/Software-Development-Backend/internal/app/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userHandling := controller.UserHandler{}

	api := app.Group("/api")

	api.Post("/register", userHandling.Register)
}
