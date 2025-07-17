package items

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *ItemHandler) {
	items := router.Group("/items")
	{
		items.POST("", handler.Create)
		items.PUT("/:id", handler.Update)
		items.DELETE("/:id", handler.Delete)
		items.GET("/:id", handler.FindByID)
		items.GET("/code/:code", handler.FindByCode)
		items.GET("", handler.FindAll)
	}
}