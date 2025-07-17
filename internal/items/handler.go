package items

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service ItemService
}

func NewItemHandler(service ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

// Create godoc
// @Summary Create a new item
// @Description Creates a new item with the provided information
// @Tags items
// @Accept json
// @Produce json
// @Param request body ItemRequest true "Item data"
// @Success 201 {object} Item "Item created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/items [post]
// @Security BearerAuth
func (h *ItemHandler) Create(c *gin.Context) {
	var request ItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	item, err := h.service.Create(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// Update godoc
// @Summary Update an item
// @Description Updates an existing item with the provided information
// @Tags items
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Param request body ItemRequest true "Item data"
// @Success 200 {object} Item "Item updated successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/items/{id} [put]
// @Security BearerAuth
func (h *ItemHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var request ItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	item, err := h.service.Update(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Delete godoc
// @Summary Delete an item
// @Description Deletes an item by its ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 204 "Item deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/items/{id} [delete]
// @Security BearerAuth
func (h *ItemHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindByID godoc
// @Summary Get an item by ID
// @Description Retrieves an item by its ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} Item "Item found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/items/{id} [get]
// @Security BearerAuth
func (h *ItemHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// FindByCode godoc
// @Summary Get an item by code
// @Description Retrieves an item by its code
// @Tags items
// @Accept json
// @Produce json
// @Param code path string true "Item code"
// @Success 200 {object} Item "Item found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/items/code/{code} [get]
// @Security BearerAuth
func (h *ItemHandler) FindByCode(c *gin.Context) {
	code := c.Param("code")
	item, err := h.service.FindByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// FindAll godoc
// @Summary Get all items
// @Description Retrieves all items
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {array} Item "List of items"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/items [get]
// @Security BearerAuth
func (h *ItemHandler) FindAll(c *gin.Context) {
	items, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}