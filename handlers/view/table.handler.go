package view_handlers

import (
	"github.com/MarcelArt/app_standard/database"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/MarcelArt/app_standard/views/components"
	"github.com/MarcelArt/app_standard/views/dev_tools"
	"github.com/gofiber/fiber/v2"
)

type TableHandler struct {
	repo repositories.ITableRepo
}

func NewTableHandler(repo repositories.ITableRepo) *TableHandler {
	return &TableHandler{
		repo: repo,
	}
}

func (h *TableHandler) Index(c *fiber.Ctx) error {
	tables, err := h.repo.GetTables()
	if err != nil {
		return utils.Render(c, dev_tools.Index([]string{}))
	}

	return utils.Render(c, dev_tools.Index(tables))
}

func (h *TableHandler) DropAll(c *fiber.Ctx) error {
	if err := database.DropDB(); err != nil {
		return utils.Render(c, components.Toast(err.Error(), "error"))
	}

	return utils.Render(c, components.Toast("Database Droped", "success"))
}
