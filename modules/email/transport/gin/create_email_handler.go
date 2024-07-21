package ginemail

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/email/biz"
	"github.com/teddlethal/web-health-check/modules/email/model"
	"github.com/teddlethal/web-health-check/modules/email/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateEmail(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var createData modelemail.EmailCreation

		if err := c.ShouldBind(&createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storageemail.NewSqlStore(db)
		business := bizemail.CreateEmailStorage(store)

		if err := business.CreateEmail(c.Request.Context(), &createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, appCommon.SimpleSuccessResponse(createData.Id))
	}
}
