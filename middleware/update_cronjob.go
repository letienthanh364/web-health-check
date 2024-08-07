package middleware

import (
	"github.com/teddlethal/web-health-check/linkchecker"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCronJobMiddleware(lc *linkchecker.LinkChecker, db *gorm.DB, deleteWebsite bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		// After request is processed
		if c.Writer.Status() == http.StatusOK {
			// Log the response body for debugging

			// For other methods, use the ID from the URL
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("Invalid website ID: %s", idStr)
				return
			}

			if c.Request.Method == http.MethodDelete && deleteWebsite {
				// Remove the cron job if the request is a DELETE
				lc.StopCronJob(id)
				log.Printf("Cron job for website ID %d removed", id)
			} else {
				// Otherwise, update the cron job
				updateCronJobBiz := bizwebsite.NewUpdateCronJobForWebsiteBiz(lc)
				go updateCronJobBiz.UpdateCronJobForWebsite(db, id)
			}
		}
	}
}
