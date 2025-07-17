package sales

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *SalesHandler) {
	router.POST("/sales/headers", handler.CreateSalesHeader)
	router.PUT("/sales/headers/:id", handler.UpdateSalesHeader)
	router.GET("/sales/headers/:id", handler.FindSalesByHeaderID)
	router.GET("/sales/headers/code/:code", handler.FindSalesByHeaderCode)
	router.GET("/sales/items/:itemCode", handler.FindSalesByItemCode)
	router.GET("/sales/customers/:customerName", handler.FindSalesByCustomerName)
	router.GET("/sales/headers", handler.FindAllSales)
	router.DELETE("/sales/headers/:id", handler.DeleteSalesByHeaderID)
	router.POST("/sales/details", handler.CreateSalesDetail)
	router.PUT("/sales/details/:id", handler.UpdateSalesDetail)
	router.GET("/sales/details/:headerID", handler.FindSalesDetailsByHeaderID)
	router.DELETE("/sales/details/:id", handler.DeleteSalesDetailByID)
	router.POST("/sales/headers/send/:id", handler.SendSalesHeader)
}
