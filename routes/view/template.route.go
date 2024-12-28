package view_routes

import (
	"github.com/MarcelArt/app_standard/database"
	view_handlers "github.com/MarcelArt/app_standard/handlers/view"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupTemplateRoutes(app *fiber.App) {
	g := app.Group("/template")

	h := view_handlers.NewTemplateHandler(repositories.NewTemplateRepo(database.GetDB()))
	g.Get("/", h.Index)
}
