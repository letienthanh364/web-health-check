package ginuser

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/user/biz"
	usermodel "github.com/teddlethal/web-health-check/modules/user/model"
	"github.com/teddlethal/web-health-check/modules/user/storage"
	"gorm.io/gorm"
	"net/http"
)

func Register(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		store := storage.NewSqlStore(db)
		md5 := appCommon.NewMd5Hash()
		business := biz.NewRegisterBiz(store, md5)

		if err := business.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(data.Id))

	}
}
