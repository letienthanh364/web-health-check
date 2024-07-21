package routes

import (
	"github.com/gin-gonic/gin"
	ginwebsite "github.com/teddlethal/web-health-check/modules/website/transport/gin"
	"gorm.io/gorm"
)

func WebsiteRoutes(router *gin.RouterGroup, db *gorm.DB, middleware func(c *gin.Context)) {
	items := router.Group("/website", middleware)
	{
		items.POST("", ginwebsite.CreateWebsite(db))
		items.GET("", ginwebsite.ListWebsite(db))
		items.GET("/:id", ginwebsite.GetWebsiteById(db))
		items.PATCH("/:id", ginwebsite.UpdateWebsite(db))
		items.DELETE("/:id", ginwebsite.DeleteWebsite(db))

		contacts := items.Group("/contact")
		{
			contacts.POST("/:id", ginwebsite.AddContactForWebsite(db))
			contacts.GET("/:id", ginwebsite.ListContactsForWebsite(db))

		}
	}

}
