package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetWebsiteById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewGetWebsiteBiz(store)

		data, err := business.GetWebsiteById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(data))
	}
}
