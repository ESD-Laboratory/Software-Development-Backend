package server

import (
	"os"

	"github.com/ESD-Laboratory/Software-Development-Backend/internal/app/routes"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	routes.SetupRoutes(app)

	err := app.Listen(":" + os.Getenv("APP_PORT"))

	if err != nil {
		panic(err)
	}

}
