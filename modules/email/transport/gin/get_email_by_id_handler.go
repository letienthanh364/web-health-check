package ginemail

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/email/biz"
	"github.com/teddlethal/web-health-check/modules/email/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetEmailById(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storageemail.NewSqlStore(db)
		business := bizemail.NewGetEmailBiz(store)

		data, err := business.GetEmailById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(data))
	}
}
