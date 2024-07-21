package routes

import (
	"github.com/gin-gonic/gin"
	gincontact "github.com/teddlethal/web-health-check/modules/contact/transport/gin"
	"gorm.io/gorm"
)

func ContactRoutes(router *gin.RouterGroup, db *gorm.DB, middleware func(c *gin.Context)) {
	items := router.Group("/contact", middleware)
	{
		items.POST("", gincontact.CreateContact(db))
		items.GET("", gincontact.ListContact(db))
		items.GET("/:id", gincontact.GetContactById(db))
		items.PATCH("/:id", gincontact.UpdateContact(db))
		items.DELETE("/:id", gincontact.DeleteContact(db))
	}
}
