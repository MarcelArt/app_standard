package routes

import (
	view_handlers "github.com/MarcelArt/app_standard/handlers/view"
	api_routes "github.com/MarcelArt/app_standard/routes/api"
	view_routes "github.com/MarcelArt/app_standard/routes/view"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())

	app.Static("/scripts", "./public/static/scripts")

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	app.Get("/", view_handlers.HelloWorldView)
	view_routes.SetupTemplateRoutes(app)

	api := app.Group("/api")
	api_routes.SetupTemplateRoutes(api)
	api_routes.SetupProcessRoutes(api)
}
