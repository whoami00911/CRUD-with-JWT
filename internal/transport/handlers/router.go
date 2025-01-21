package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (har *HandlerAssetsResponse) InitRoutes() *gin.Engine {
	router := gin.Default()
	abuseipGroup := router.Group("/Abuseip")
	{
		abuseipGroup.POST("/", har.CreateHandler)
		abuseipGroup.PUT("/", har.UpdateHandler)
		abuseipGroup.GET("/", har.GetAllHandler)
		abuseipGroup.DELETE("/", har.DeleteAllHandler)

		// Обработка маршрутов с IP
		abuseipGroup.GET("/:ip", har.GetHandler)
		abuseipGroup.DELETE("/:ip", har.DeleteHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
