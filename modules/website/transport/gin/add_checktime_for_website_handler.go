package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AddCheckTimeForWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var addCheckTimeData modelwebsite.WebsiteCheckTimeCreation

		if err := c.ShouldBind(&addCheckTimeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := addCheckTimeData.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		websiteStorage := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewAddCheckTimeForWebsiteBiz(websiteStorage)

		if err := business.AddCheckTimeForWebsite(c.Request.Context(), id, &addCheckTimeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
