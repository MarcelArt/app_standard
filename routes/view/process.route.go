package view_routes

import (
	"github.com/MarcelArt/app_standard/database"
	view_handlers "github.com/MarcelArt/app_standard/handlers/view"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupProcessRoutes(app *fiber.App) {
	g := app.Group("/process")

	h := view_handlers.NewProcessHandler(repositories.NewProcessRepo(database.GetDB()))
	g.Get("/", h.Index)
	g.Get("/create", h.CreateView)

	g.Post("/create", h.Create)
}
