package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/linkchecker"
	"github.com/teddlethal/web-health-check/middleware"
	ginwebsite "github.com/teddlethal/web-health-check/modules/website/transport/gin"
	"gorm.io/gorm"
)

func WebsiteRoutes(router *gin.RouterGroup, db *gorm.DB, authMiddleware func(c *gin.Context), lc *linkchecker.LinkChecker) {
	items := router.Group("/website", authMiddleware)
	{
		items.POST("", ginwebsite.CreateWebsite(db, lc))
		items.GET("", ginwebsite.ListWebsite(db))
		items.GET("/:id", ginwebsite.GetWebsiteById(db))
		items.PATCH("/:id", ginwebsite.UpdateWebsite(db), middleware.UpdateCronJobMiddleware(lc, db, false))
		items.DELETE("/:id", ginwebsite.DeleteWebsite(db), middleware.UpdateCronJobMiddleware(lc, db, true))

		contacts := items.Group("/contact")
		{
			contacts.POST("/:id", ginwebsite.AddContactForWebsite(db), middleware.UpdateCronJobMiddleware(lc, db, false))
			contacts.GET("/:id", ginwebsite.ListContactsForWebsite(db))
			contacts.DELETE("/:id", ginwebsite.DeleteContactForWebsite(db), middleware.UpdateCronJobMiddleware(lc, db, false))
		}

		checktimes := items.Group("/check-time")
		{
			checktimes.POST("/:id", ginwebsite.AddCheckTimeForWebsite(db), middleware.UpdateCronJobMiddleware(lc, db, false))
			checktimes.GET("/:id", ginwebsite.ListCheckTimesForWebsite(db))
			checktimes.DELETE("/:id", ginwebsite.DeleteCheckTimeForWebsite(db), middleware.UpdateCronJobMiddleware(lc, db, false))

		}
	}
}
