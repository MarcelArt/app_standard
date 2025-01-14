package api_handlers

import (
	"github.com/MarcelArt/app_standard/models"
	"github.com/MarcelArt/app_standard/repositories"
	"github.com/MarcelArt/app_standard/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	BaseCrudHandler[models.User, models.UserDTO, models.UserPage]
	repo repositories.IUserRepo
}

func NewUserHandler(repo repositories.IUserRepo) *UserHandler {
	return &UserHandler{
		BaseCrudHandler: BaseCrudHandler[models.User, models.UserDTO, models.UserPage]{
			repo: repo,
		},
		repo: repo,
	}
}

// Create creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param User body models.UserDTO true "User data"
// @Success 201 {object} models.UserDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user models.UserDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user.Salt = utils.RandString(10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+user.Salt), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	user.Password = string(hashedPassword)

	id, err := h.repo.Create(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"ID": id})
}

// Read retrieves a list of users
// @Summary Get a list of users
// @Description Get a list of users
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.UserPage
// @Router /user [get]
func (h *UserHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing user
// @Summary Update an existing user
// @Description Update an existing user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param User body models.UserDTO true "User data"
// @Success 200 {object} models.UserDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing user
// @Summary Delete an existing user
// @Description Delete an existing user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} string
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a user by ID
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} string
// @Router /user/{id} [get]
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}

// Login is a function to login
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param input body models.LoginInput true "Login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /user/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var user models.LoginInput
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	userDB, err := h.repo.GetByUsernameOrEmail(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password+userDB.Salt))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(userDB, user.IsRemeber)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
