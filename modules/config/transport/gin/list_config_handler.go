package ginconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/biz"
	"github.com/teddlethal/web-health-check/modules/config/model"
	"github.com/teddlethal/web-health-check/modules/config/storage"
	"gorm.io/gorm"
	"net/http"
)

func ListConfig(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelconfig.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storage.NewSqlStore(db)
		business := biz.NewListConfigBiz(store)

		res, err := business.ListConfig(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, queryString.Filter, queryString.Paging))

	}
}
