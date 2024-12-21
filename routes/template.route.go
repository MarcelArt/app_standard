package routes

import (
	"github.com/MarcelArt/app_standard/database"
	"github.com/MarcelArt/app_standard/handlers"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupTemplateRoutes(app *fiber.App) {
	h := handlers.NewTemplateHandler(repositories.NewTemplateRepo(database.GetDB()))

	g := app.Group("/template")

	// Read retrieves a list of templates
	// @Summary Get a list of templates
	// @Description Get a list of templates
	// @Tags templates
	// @Accept json
	// @Produce json
	// @Success 200 {array} models.TemplatePage
	// @Router /template [get]
	g.Get("/", h.Read)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
	g.Put("/:id", h.Update)
	g.Delete("/:id", h.Delete)
}
