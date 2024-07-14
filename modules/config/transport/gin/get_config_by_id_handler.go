package ginconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/biz"
	"github.com/teddlethal/web-health-check/modules/config/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetConfigById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storageconfig.NewSqlStore(db)
		business := biz.NewGetConfigBiz(store)

		data, err := business.GetConfigById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(data))
	}
}
