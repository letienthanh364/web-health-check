package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/checker"
	"net/http"
)

func CheckerRoutes(router *gin.RouterGroup) {
	checking := router.Group("/check-link")
	{
		checking.POST("", func(c *gin.Context) {
			var json struct {
				URL string `json:"url" binding:"required"`
			}

			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			isDead := checker.CheckLink(json.URL)
			status := "alive"
			if isDead {
				status = "dead"
			}

			c.JSON(http.StatusOK, gin.H{
				"status": status,
			})
		})

	}
}
