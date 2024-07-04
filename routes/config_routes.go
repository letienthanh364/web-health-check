package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/modules/config/transport/gin"
	"gorm.io/gorm"
)

func ConfigRoutes(router *gin.RouterGroup, db *gorm.DB) {
	items := router.Group("/config")
	{
		items.POST("", ginconfig.CreateConfig(db))
		items.GET("", ginconfig.ListConfig(db))
		items.GET("/:configId", ginconfig.GetConfigById(db))
		items.PATCH("/:configId", ginconfig.UpdateConfig(db))
		items.DELETE("/:configId", ginconfig.DeleteConfig(db))
	}
}
