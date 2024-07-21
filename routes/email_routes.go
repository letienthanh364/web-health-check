package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/modules/email/transport/gin"
	"gorm.io/gorm"
)

func EmailRoutes(router *gin.RouterGroup, db *gorm.DB, middleware func(c *gin.Context)) {
	items := router.Group("/email", middleware)
	{
		items.POST("", ginemail.CreateEmail(db))
		items.GET("", ginemail.ListEmail(db))
		items.GET("/:id", ginemail.GetEmailById(db))
		items.PATCH("/:id", ginemail.UpdateEmail(db))
		items.DELETE("/:id", ginemail.DeleteEmail(db))
	}
}
