package router

import (
	"github.com/dickanirwansyah/blogspot/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)
}
