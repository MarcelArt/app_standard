package routes

import (
	"github.com/MarcelArt/app_standard/database"
	api_handlers "github.com/MarcelArt/app_standard/handlers/api"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupTemplateRoutes(api fiber.Router) {
	h := api_handlers.NewTemplateHandler(repositories.NewTemplateRepo(database.GetDB()))

	g := api.Group("/template")

	g.Get("/", h.Read)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
	g.Put("/:id", h.Update)
	g.Delete("/:id", h.Delete)
}
