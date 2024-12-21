package repositories

import (
	"github.com/MarcelArt/app_standard/models"
	"gorm.io/gorm"
)

const processPageQuery = `
	SELECT 
		p.id id, 
		p.template_id template_id,
		t.name template_name
	FROM processes p
	join templates t on p.template_id = t.id	
`

type IProcessRepo interface {
	IBaseCrudRepo[models.Process, models.ProcessDTO, models.ProcessPage]
}

type ProcessRepo struct {
	BaseCrudRepo[models.Process, models.ProcessDTO, models.ProcessPage]
}

func NewProcessRepo(db *gorm.DB) *ProcessRepo {
	return &ProcessRepo{
		BaseCrudRepo: BaseCrudRepo[models.Process, models.ProcessDTO, models.ProcessPage]{
			db:        db,
			pageQuery: processPageQuery,
		},
	}
}
