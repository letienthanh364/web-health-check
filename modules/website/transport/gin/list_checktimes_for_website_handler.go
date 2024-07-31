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

func ListCheckTimesForWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelwebsite.WebsiteCheckTimeFilter
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		websiteStorage := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewListCheckTimesForWebsiteBiz(websiteStorage)

		res, err := business.ListCheckTimesForWebsite(c.Request.Context(), id, &queryString.WebsiteCheckTimeFilter, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(res))
	}
}
