package scaffold

import (
	"os"
	"strings"
)

const routeTemplate = `
package routes

import (
	"${moduleName}/database"
	api_handlers "github.com/MarcelArt/app_standard/handlers/api"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

func Setup${modelName}Routes(api fiber.Router) {
	h := api_handlers.New${modelName}Handler(repositories.New${modelName}Repo(database.GetDB()))

	g := api.Group("/${handlerRoute}")
	g.Get("/", h.Read)
	g.Get("/:id", h.GetByID)
	g.Post("/", h.Create)
	g.Put("/:id", h.Update)
	g.Delete("/:id", h.Delete)
}
`

func ScaffoldRoute(modelName string, handlerRoute string) {
	newRoute := strings.ReplaceAll(routeTemplate, "${modelName}", modelName)
	newRoute = strings.ReplaceAll(newRoute, "${handlerRoute}", handlerRoute)
	newRoute = strings.ReplaceAll(newRoute, "${moduleName}", moduleName)

	filename := "routes/" + ToSeparateByCharLowered(modelName, '_') + ".route.go"
	if err := os.WriteFile(filename, []byte(newRoute), 0644); err != nil {
		panic("Failed writing route file")
	}
}
