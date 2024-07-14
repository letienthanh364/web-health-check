package ginuser

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/component/tokenprovider"
	"github.com/teddlethal/web-health-check/modules/user/biz"
	usermodel "github.com/teddlethal/web-health-check/modules/user/model"
	"github.com/teddlethal/web-health-check/modules/user/storage"
	"gorm.io/gorm"
	"net/http"
)

func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		store := storageuser.NewSqlStore(db)
		md5 := appCommon.NewMd5Hash()
		business := biz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(account))
	}

}
