package view_handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/MarcelArt/app_standard/views/base_crud"
	"github.com/gofiber/fiber/v2"
)

type BaseCrudHandler[TModel any, TDto models.Formable, TPage models.IViewable] struct {
	repo      repositories.IBaseCrudRepo[TModel, TDto, TPage]
	PageName  string
	PageDesc  string
	PageRoute string
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Index(c *fiber.Ctx) error {
	var dest []TPage

	page := h.repo.Read(c, dest)

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

	return utils.Render(c, base_crud.Index(response, h.PageRoute))
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) CreateView(c *fiber.Ctx) error {
	var dto TDto

	inputForm := dto.ToView()

	return utils.Render(c, base_crud.Create(inputForm, h.PageRoute))
}

func (h *BaseCrudHandler[TModel, TDto, TPage]) Create(c *fiber.Ctx) error {
	var dto TDto

	inputForm := dto.ToView()
	for key, value := range inputForm {
		inputForm[key] = c.FormValue(key, value)
	}

	input, err := dto.FromView(inputForm)
	if err != nil {
		return c.Redirect(h.PageRoute)
	}

	dto, ok := input.(TDto)
	if !ok {
		return c.Redirect(h.PageRoute)
	}

	h.repo.Create(dto)

	return c.Redirect(h.PageRoute)
}
