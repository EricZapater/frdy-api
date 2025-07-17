package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Create godoc
// @Summary Register a new user
// @Description Register a new user with the provided information (Public route)
// @Tags auth
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User registration data"
// @Success 201 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (h *UserHandler) Create(c *gin.Context) {	
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Create(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Update godoc
// @Summary Update a user
// @Description Update user information by ID (Protected route)
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "User update data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users/{id} [put]
// @Security BearerAuth
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var request UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Update(c.Request.Context(), id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Delete a user
// @Description Delete a user by ID (Protected route)
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/users/{id} [delete]
// @Security BearerAuth
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.userService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password for a user (Protected route)
// @Tags users
// @Accept json
// @Produce json
// @Param password body ChangePasswordRequest true "Password change data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users/change-password [post]
// @Security BearerAuth
func (h *UserHandler) ChangePassword(c *gin.Context) {
	var request ChangePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.ChangePassword(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetByUsername godoc
// @Summary Get user by username
// @Description Retrieve a user by their username (Protected route)
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} UserResponse
// @Failure 500 {object} map[string]string
// @Router /api/users/username/{username} [get]
// @Security BearerAuth
func(h *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userService.FindByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetByID godoc
// @Summary Get user by ID
// @Description Retrieve a user by their ID (Protected route)
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 500 {object} map[string]string
// @Router /api/users/{id} [get]
// @Security BearerAuth
func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.FindByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAll godoc
// @Summary Get all users
// @Description Retrieve all users (Protected route)
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} UserResponse
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
// @Security BearerAuth
func(h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.userService.FindAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}