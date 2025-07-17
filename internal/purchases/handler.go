package purchases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchasesHandler struct {
	service PurchaseService
}

func NewPurchasesHandler(service PurchaseService) *PurchasesHandler {
	return &PurchasesHandler{service: service}
}

// CreatePurchaseHeader godoc
// @Summary Create a new purchase header
// @Description Creates a new purchase header with the provided information
// @Tags purchases
// @Accept json
// @Produce json
// @Param request body PurchaseHeaderRequest true "Purchase header data"
// @Success 201 {object} PurchaseHeader "Purchase header created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/headers [post]
// @Security BearerAuth
func (h *PurchasesHandler) CreatePurchaseHeader(c *gin.Context) {
	var request PurchaseHeaderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	header, err := h.service.CreatePurchaseHeader(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, header)
}

// UpdatePurchaseHeader godoc
// @Summary Update a purchase header
// @Description Updates an existing purchase header with the provided information
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase header ID"
// @Param request body PurchaseHeaderRequest true "Purchase header data"
// @Success 200 {object} PurchaseHeader "Purchase header updated successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/headers/{id} [put]
// @Security BearerAuth
func (h *PurchasesHandler) UpdatePurchaseHeader(c *gin.Context) {
	id := c.Param("id")
	var request PurchaseHeaderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	header, err := h.service.UpdatePurchaseHeader(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

// FindPurchaseByID godoc
// @Summary Get a purchase header by ID
// @Description Retrieves a purchase header by its ID
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase header ID"
// @Success 200 {object} PurchaseHeader "Purchase header found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/headers/{id} [get]
// @Security BearerAuth
func (h *PurchasesHandler) FindPurchaseByID(c *gin.Context) {
	id := c.Param("id")
	header, err := h.service.FindPurchaseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

// FindAllPurchases godoc
// @Summary Get all purchase headers
// @Description Retrieves all purchase headers
// @Tags purchases
// @Accept json
// @Produce json
// @Success 200 {array} PurchaseHeader "List of purchase headers"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/headers [get]
// @Security BearerAuth
func (h *PurchasesHandler) FindAllPurchases(c *gin.Context) {
	headers, err := h.service.FindAllPurchases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, headers)
}

// DeletePurchaseByID godoc
// @Summary Delete a purchase header by ID
// @Description Deletes a purchase header by its ID
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase header ID"
// @Success 204 "Purchase header deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/headers/{id} [delete]
// @Security BearerAuth
func (h *PurchasesHandler) DeletePurchaseByID(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeletePurchaseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreatePurchaseDetail godoc
// @Summary Create a new purchase detail
// @Description Creates a new purchase detail with the provided information
// @Tags purchases
// @Accept json
// @Produce json
// @Param request body PurchaseDetailRequest true "Purchase detail data"
// @Success 201 {object} PurchaseDetail "Purchase detail created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/details [post]
// @Security BearerAuth
func (h *PurchasesHandler) CreatePurchaseDetail(c *gin.Context) {
	var request PurchaseDetailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := h.service.CreatePurchaseDetail(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, detail)
}

// UpdatePurchaseDetail godoc
// @Summary Update a purchase detail
// @Description Updates an existing purchase detail with the provided information
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase detail ID"
// @Param request body PurchaseDetailRequest true "Purchase detail data"
// @Success 200 {object} PurchaseDetail "Purchase detail updated successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/details/{id} [put]
// @Security BearerAuth
func (h *PurchasesHandler) UpdatePurchaseDetail(c *gin.Context) {
	id := c.Param("id")
	var request PurchaseDetailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := h.service.UpdatePurchaseDetail(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

// FindDetailsByPurchaseID godoc
// @Summary Get purchase details by header ID
// @Description Retrieves all purchase details for a specific purchase header
// @Tags purchases
// @Accept json
// @Produce json
// @Param header_id path string true "Purchase header ID"
// @Success 200 {array} PurchaseDetail "List of purchase details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/details/{header_id} [get]
// @Security BearerAuth
func (h *PurchasesHandler) FindDetailsByPurchaseID(c *gin.Context) {
	headerID := c.Param("header_id")
	details, err := h.service.FindDetailsByPurchaseID(headerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}

// DeletePurchaseDetailByID godoc
// @Summary Delete a purchase detail by ID
// @Description Deletes a purchase detail by its ID
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path string true "Purchase detail ID"
// @Success 204 "Purchase detail deleted successfully"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /purchases/details/{id} [delete]
// @Security BearerAuth
func (h *PurchasesHandler) DeletePurchaseDetailByID(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeletePurchaseDetailByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}