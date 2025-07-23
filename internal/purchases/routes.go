package purchases

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *PurchasesHandler) {
	// Headers
	router.POST("/purchases/headers", handler.CreatePurchaseHeader)
	router.PUT("/purchases/headers/:id", handler.UpdatePurchaseHeader)
	router.GET("/purchases/headers/:id", handler.FindPurchaseByID)
	router.GET("/purchases/headers", handler.FindAllPurchases)
	router.DELETE("/purchases/headers/:id", handler.DeletePurchaseByID)
	router.GET("/purchases/headers/receive/:id", handler.ReceivePurchaseHeader)

	// Details
	router.POST("/purchases/details", handler.CreatePurchaseDetail)
	router.PUT("/purchases/details/:id", handler.UpdatePurchaseDetail)
	router.GET("/purchases/details/:header_id", handler.FindDetailsByPurchaseID)
	router.DELETE("/purchases/details/:id", handler.DeletePurchaseDetailByID)
}
