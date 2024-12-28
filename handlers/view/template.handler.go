package view_handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
)

type TemplateHandler struct {
	BaseCrudHandler[models.Template, models.TemplateDTO, models.TemplatePage]
}

func NewTemplateHandler(repo repositories.ITemplateRepo) *TemplateHandler {
	return &TemplateHandler{
		BaseCrudHandler: BaseCrudHandler[models.Template, models.TemplateDTO, models.TemplatePage]{
			repo:     repo,
			PageName: "Template",
			PageDesc: "Template Description",
		},
	}
}
