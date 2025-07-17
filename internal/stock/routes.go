package stock

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *StockHandler) {
	router.GET("/stock", handler.GetAllStocks)
	router.GET("/stock/:item_id", handler.GetStockByItemID)
	router.PUT("/stock/:item_id", handler.UpdateStockQuantity)
}
