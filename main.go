package main

import (
	"fmt"
	"log"

	"github.com/MarcelArt/app_standard/config"
	_ "github.com/MarcelArt/app_standard/config"
	"github.com/MarcelArt/app_standard/database"
	_ "github.com/MarcelArt/app_standard/docs"
	"github.com/MarcelArt/app_standard/routes"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	serve()
}

func serve() {
	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Printf("Listening to http://localhost:%s", config.Env.PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.PORT)))
}
