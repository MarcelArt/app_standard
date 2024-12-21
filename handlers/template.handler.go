package handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

type TemplateHandler struct {
	BaseCrudHandler[models.Template, models.TemplateDTO, models.TemplatePage]
}

func NewTemplateHandler(repo repositories.ITemplateRepo) *TemplateHandler {
	return &TemplateHandler{
		BaseCrudHandler: BaseCrudHandler[models.Template, models.TemplateDTO, models.TemplatePage]{
			repo: repo,
		},
	}
}

// Create creates a new template
// @Summary Create a new template
// @Description Create a new template
// @Tags Template
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param template body models.TemplateDTO true "Template data"
// @Success 201 {object} models.TemplateDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /template [post]
func (h *TemplateHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of templates
// @Summary Get a list of templates
// @Description Get a list of templates
// @Tags Template
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.TemplatePage
// @Router /template [get]
func (h *TemplateHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing template
// @Summary Update an existing template
// @Description Update an existing template
// @Tags Template
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Template ID"
// @Param template body models.TemplateDTO true "Template data"
// @Success 200 {object} models.TemplateDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /template/{id} [put]
func (h *TemplateHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing template
// @Summary Delete an existing template
// @Description Delete an existing template
// @Tags Template
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Template ID"
// @Success 200 {object} models.Template
// @Failure 500 {object} string
// @Router /template/{id} [delete]
func (h *TemplateHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a template by ID
// @Summary Get a template by ID
// @Description Get a template by ID
// @Tags Template
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Template ID"
// @Success 200 {object} models.Template
// @Failure 500 {object} string
// @Router /template/{id} [get]
func (h *TemplateHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
