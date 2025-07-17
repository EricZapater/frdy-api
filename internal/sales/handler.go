package sales

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SalesHandler struct {
	service SalesService
}

func NewSalesHandler(service SalesService) *SalesHandler {
	return &SalesHandler{service: service}
}

// CreateSalesHeader godoc
// @Summary Create a new sales header
// @Description Create a new sales header with the provided information (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Param header body SalesHeaderRequest true "Sales header data"
// @Success 201 {object} SalesHeader
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers [post]
// @Security BearerAuth
func (h *SalesHandler) CreateSalesHeader(c *gin.Context) {
	var request SalesHeaderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	header, err := h.service.CreateSalesHeader(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, header)
}

// UpdateSalesHeader godoc
// @Summary Update a sales header
// @Description Update sales header information by ID (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Param id path string true "Sales Header ID"
// @Param header body SalesHeaderRequest true "Sales header update data"
// @Success 200 {object} SalesHeader
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers/{id} [put]
// @Security BearerAuth
func (h *SalesHandler) UpdateSalesHeader(c *gin.Context) {
	id := c.Param("id")
	var request SalesHeaderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	header, err := h.service.UpdateSalesHeader(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

// FindSalesByHeaderID godoc
// @Summary Get sales by header ID
// @Description Retrieve sales information by header ID (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Param id path string true "Sales Header ID"
// @Success 200 {object} SalesHeader
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers/{id} [get]
// @Security BearerAuth
func (h *SalesHandler) FindSalesByHeaderID(c *gin.Context) {
	id := c.Param("id")
	header, err := h.service.FindSalesByHeaderID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

// FindSalesByHeaderCode godoc
// @Summary Get sales by header code
// @Description Retrieve sales information by header code (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Param code path string true "Sales Header Code"
// @Success 200 {object} SalesHeader
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers/code/{code} [get]
// @Security BearerAuth
func (h *SalesHandler) FindSalesByHeaderCode(c *gin.Context) {
	code := c.Param("code")
	header, err := h.service.FindSalesByHeaderCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

// FindSalesByItemCode godoc
// @Summary Get sales by item code
// @Description Retrieve sales information by item code (Protected route)
// @Tags sales-queries
// @Accept json
// @Produce json
// @Param itemCode path string true "Item Code"
// @Success 200 {array} SalesHeader
// @Failure 500 {object} map[string]string
// @Router /api/sales/items/{itemCode} [get]
// @Security BearerAuth
func (h *SalesHandler) FindSalesByItemCode(c *gin.Context) {
	itemCode := c.Param("itemCode")
	headers, err := h.service.FindSalesByItemCode(itemCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, headers)
}

// FindSalesByCustomerName godoc
// @Summary Get sales by customer name
// @Description Retrieve sales information by customer name (Protected route)
// @Tags sales-queries
// @Accept json
// @Produce json
// @Param customerName path string true "Customer Name"
// @Success 200 {array} SalesHeader
// @Failure 500 {object} map[string]string
// @Router /api/sales/customers/{customerName} [get]
// @Security BearerAuth
func (h *SalesHandler) FindSalesByCustomerName(c *gin.Context) {
	customerName := c.Param("customerName")
	headers, err := h.service.FindSalesByCustomerName(customerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, headers)
}

// FindAllSales godoc
// @Summary Get all sales
// @Description Retrieve all sales headers (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Success 200 {array} SalesHeader
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers [get]
// @Security BearerAuth
func (h *SalesHandler) FindAllSales(c *gin.Context) {
	headers, err := h.service.FindAllSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, headers)
}

// DeleteSalesByHeaderID godoc
// @Summary Delete sales by header ID
// @Description Delete sales information by header ID (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Param id path string true "Sales Header ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers/{id} [delete]
// @Security BearerAuth
func (h *SalesHandler) DeleteSalesByHeaderID(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteSalesByHeaderID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// SendSalesHeader godoc
// @Summary Send sales header
// @Description Send/confirm a sales header by ID (Protected route)
// @Tags sales-headers
// @Accept json
// @Produce json
// @Param id path string true "Sales Header ID"
// @Success 200 {object} SalesHeader
// @Failure 500 {object} map[string]string
// @Router /api/sales/headers/send/{id} [post]
// @Security BearerAuth
func (h *SalesHandler) SendSalesHeader(c *gin.Context) {
	id := c.Param("id")
	header, err := h.service.SendSalesHeader(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, header)
}

// CreateSalesDetail godoc
// @Summary Create a new sales detail
// @Description Create a new sales detail with the provided information (Protected route)
// @Tags sales-details
// @Accept json
// @Produce json
// @Param detail body SalesDetailRequest true "Sales detail data"
// @Success 201 {object} SalesDetail
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/sales/details [post]
// @Security BearerAuth
func (h *SalesHandler) CreateSalesDetail(c *gin.Context) {
	var request SalesDetailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	detail, err := h.service.CreateSalesDetail(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, detail)
}

// UpdateSalesDetail godoc
// @Summary Update a sales detail
// @Description Update sales detail information by ID (Protected route)
// @Tags sales-details
// @Accept json
// @Produce json
// @Param id path string true "Sales Detail ID"
// @Param detail body SalesDetailRequest true "Sales detail update data"
// @Success 200 {object} SalesDetail
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/sales/details/{id} [put]
// @Security BearerAuth
func (h *SalesHandler) UpdateSalesDetail(c *gin.Context) {
	id := c.Param("id")
	var request SalesDetailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	detail, err := h.service.UpdateSalesDetail(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

// DeleteSalesDetailByID godoc
// @Summary Delete sales detail by ID
// @Description Delete sales detail information by ID (Protected route)
// @Tags sales-details
// @Accept json
// @Produce json
// @Param id path string true "Sales Detail ID"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/sales/details/{id} [delete]
// @Security BearerAuth
func (h *SalesHandler) DeleteSalesDetailByID(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteSalesDetailByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// FindSalesDetailsByHeaderID godoc
// @Summary Get sales details by header ID
// @Description Retrieve all sales details for a specific header ID (Protected route)
// @Tags sales-details
// @Accept json
// @Produce json
// @Param headerID path string true "Sales Header ID"
// @Success 200 {array} SalesDetail
// @Failure 500 {object} map[string]string
// @Router /api/sales/details/{headerID} [get]
// @Security BearerAuth
func (h *SalesHandler) FindSalesDetailsByHeaderID(c *gin.Context) {
	headerID := c.Param("headerID")
	details, err := h.service.FindSalesDetailsByHeaderID(headerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}