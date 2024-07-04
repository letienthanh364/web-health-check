package ginconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/biz"
	configmodel "github.com/teddlethal/web-health-check/modules/config/model"
	"github.com/teddlethal/web-health-check/modules/config/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateConfig(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var createData configmodel.ConfigCreation

		if err := c.ShouldBind(&createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSqlStore(db)
		business := biz.CreateConfigStorage(store)

		if err := business.CreateConfig(c.Request.Context(), &createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, appCommon.SimpleSuccessResponse(createData.Id))
	}
}
