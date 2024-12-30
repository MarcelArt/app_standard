package view_handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
)

type ProcessHandler struct {
	BaseCrudHandler[models.Process, models.ProcessDTO, models.ProcessPage]
}

func NewProcessHandler(repo repositories.IProcessRepo) *ProcessHandler {
	return &ProcessHandler{
		BaseCrudHandler: BaseCrudHandler[models.Process, models.ProcessDTO, models.ProcessPage]{
			repo:      repo,
			PageName:  "Process",
			PageDesc:  "Process Description",
			PageRoute: "/process",
		},
	}
}
