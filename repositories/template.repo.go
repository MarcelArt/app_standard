package repositories

import (
	"github.com/MarcelArt/app_standard/models"
	"gorm.io/gorm"
)

const templatePageQuery = `
	SELECT 
		t.id id, 
		t.name name
	FROM templates t
`

type ITemplateRepo interface {
	IBaseCrudRepo[models.Template, models.TemplateDTO, models.TemplatePage]
}

type TemplateRepo struct {
	BaseCrudRepo[models.Template, models.TemplateDTO, models.TemplatePage]
}

func NewTemplateRepo(db *gorm.DB) *TemplateRepo {
	return &TemplateRepo{
		BaseCrudRepo: BaseCrudRepo[models.Template, models.TemplateDTO, models.TemplatePage]{
			db:        db,
			pageQuery: templatePageQuery,
		},
	}
}
