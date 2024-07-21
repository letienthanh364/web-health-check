package gincontact

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/contact/biz"
	"github.com/teddlethal/web-health-check/modules/contact/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetContactById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storagecontact.NewSqlStore(db)
		business := bizecontact.NewGetContactBiz(store)

		data, err := business.GetContactById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(data))
	}
}
