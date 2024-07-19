package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"

	"gorm.io/gorm"
	"net/http"
)

func ListWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelwebsite.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewListWebsiteBiz(store)

		res, err := business.ListWebsite(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, queryString.Filter, queryString.Paging))

	}
}
