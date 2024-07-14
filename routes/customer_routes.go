package routes

import (
	"github.com/gin-gonic/gin"
	ginconfig "github.com/teddlethal/web-health-check/modules/config/transport/gin"
	gincustomer "github.com/teddlethal/web-health-check/modules/customer/transport/gin"
	"gorm.io/gorm"
)

func CustomerRoutes(router *gin.RouterGroup, db *gorm.DB) {
	items := router.Group("/customer")
	{
		items.POST("", gincustomer.CreateCustomer(db))
		items.GET("", ginconfig.ListConfig(db))
		items.GET("/:configId", ginconfig.GetConfigById(db))
		items.PATCH("/:configId", ginconfig.UpdateConfig(db))
		items.DELETE("/:configId", ginconfig.DeleteConfig(db))
	}
}
