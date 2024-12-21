package handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/gofiber/fiber/v2"
)

type ProcessHandler struct {
	BaseCrudHandler[models.Process, models.ProcessDTO, models.ProcessPage]
}

func NewProcessHandler(repo repositories.IProcessRepo) *ProcessHandler {
	return &ProcessHandler{
		BaseCrudHandler: BaseCrudHandler[models.Process, models.ProcessDTO, models.ProcessPage]{
			repo: repo,
		},
	}
}

// Create creates a new process
// @Summary Create a new process
// @Description Create a new process
// @Tags Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param process body models.ProcessDTO true "Process data"
// @Success 201 {object} models.ProcessDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /process [post]
func (h *ProcessHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of processes
// @Summary Get a list of processes
// @Description Get a list of processes
// @Tags Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.ProcessPage
// @Router /process [get]
func (h *ProcessHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing process
// @Summary Update an existing process
// @Description Update an existing process
// @Tags Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Process ID"
// @Param process body models.ProcessDTO true "Process data"
// @Success 200 {object} models.ProcessDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /process/{id} [put]
func (h *ProcessHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing process
// @Summary Delete an existing process
// @Description Delete an existing process
// @Tags Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Process ID"
// @Success 200 {object} models.Process
// @Failure 500 {object} string
// @Router /process/{id} [delete]
func (h *ProcessHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a process by ID
// @Summary Get a process by ID
// @Description Get a process by ID
// @Tags Process
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Process ID"
// @Success 200 {object} models.Process
// @Failure 500 {object} string
// @Router /process/{id} [get]
func (h *ProcessHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
