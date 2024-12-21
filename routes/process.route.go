package routes

import (
	"github.com/MarcelArt/app_standard/database"
	"github.com/MarcelArt/app_standard/handlers"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupProcessRoutes(app *fiber.App) {
	h := handlers.NewProcessHandler(repositories.NewProcessRepo(database.GetDB()))

	g := app.Group("/process")
	g.Get("/", h.Read)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
	g.Put("/:id", h.Update)
	g.Delete("/:id", h.Delete)
}
