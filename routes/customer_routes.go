package routes

import (
	"github.com/gin-gonic/gin"
	gincustomer "github.com/teddlethal/web-health-check/modules/customer/transport/gin"
	"gorm.io/gorm"
)

func CustomerRoutes(router *gin.RouterGroup, db *gorm.DB, middleware func(c *gin.Context)) {
	items := router.Group("/customer", middleware)
	{
		items.POST("", gincustomer.CreateCustomer(db))
		items.GET("", gincustomer.ListCustomer(db))
		items.GET("/:id", gincustomer.FindCustomer(db))
		items.PATCH("/:id", gincustomer.UpdateCustomer(db))
		items.DELETE("/:id", gincustomer.DeleteCustomer(db))
	}
}
