package routes

import (
	"time"

	"github.com/MarcelArt/app_standard/database"
	"github.com/MarcelArt/app_standard/middlewares"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	app.Get("/metrics", monitor.New())

	authMiddleware := middlewares.NewAuthMiddleware(repositories.NewUserRepo(database.GetDB()))

	api := app.Group("/api")
	SetupUserRoutes(api, authMiddleware)
	SetupAuthorizedDeviceRoutes(api, authMiddleware)
}
