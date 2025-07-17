package stock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	service StockService
}

func NewStockHandler(service StockService) *StockHandler {
	return &StockHandler{service: service}
}

// UpdateStockQuantityRequest represents the request body for updating stock quantity
type UpdateStockQuantityRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}

// GetStockByItemID godoc
// @Summary Get stock by item ID
// @Description Retrieve stock information for a specific item by its ID (Protected route)
// @Tags stock
// @Accept json
// @Produce json
// @Param item_id path string true "Item ID"
// @Success 200 {object} Stock
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/stock/{item_id} [get]
// @Security BearerAuth
func (h *StockHandler) GetStockByItemID(c *gin.Context) {
	itemID := c.Param("item_id")
	stock, err := h.service.GetStockByItemID(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if (stock == Stock{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}

// UpdateStockQuantity godoc
// @Summary Update stock quantity
// @Description Update the quantity of stock for a specific item (Protected route)
// @Tags stock
// @Accept json
// @Produce json
// @Param item_id path string true "Item ID"
// @Param quantity body UpdateStockQuantityRequest true "Stock quantity update data"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/stock/{item_id} [put]
// @Security BearerAuth
func (h *StockHandler) UpdateStockQuantity(c *gin.Context) {
	itemID := c.Param("item_id")
	var request UpdateStockQuantityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.service.UpdateStockQuantity(itemID, request.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetAllStocks godoc
// @Summary Get all stocks
// @Description Retrieve all stock information (Protected route)
// @Tags stock
// @Accept json
// @Produce json
// @Success 200 {array} Stock
// @Failure 500 {object} map[string]string
// @Router /api/stock [get]
// @Security BearerAuth
func (h *StockHandler) GetAllStocks(c *gin.Context) {
	stocks, err := h.service.GetAllStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}