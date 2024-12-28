package view_handlers

import (
	"log"

	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/MarcelArt/app_standard/views/base_crud"
	"github.com/gofiber/fiber/v2"
)

type BaseCrudHandler[TModel any, TDto models.IDTO, TPage models.IViewable] struct {
	repo     repositories.IBaseCrudRepo[TModel, TDto, TPage]
	PageName string
	PageDesc string
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Index(c *fiber.Ctx) error {
	var dest []TPage

	page := h.repo.Read(c, dest)

	log.Println(page.Items)

	datas, ok := page.Items.(*[]TPage)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to convert items"})
	}

	items := utils.SliceMap(*datas, func(currentValue TPage) models.View {
		return currentValue.ToView()
	})

	var columns []string
	for key := range items[0] {
		columns = append(columns, key)
	}

	response := models.PageView{
		Page:        page,
		Items:       items,
		Columns:     columns,
		Title:       h.PageName,
		Description: h.PageDesc,
	}

	return utils.Render(c, base_crud.Index(response))
}
