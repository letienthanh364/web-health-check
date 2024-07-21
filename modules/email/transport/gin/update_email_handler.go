package ginemail

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizemail "github.com/teddlethal/web-health-check/modules/email/biz"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
	storageemail "github.com/teddlethal/web-health-check/modules/email/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateEmail(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var updateData modelemail.EmailUpdate

		if err := c.ShouldBind(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storageemail.NewSqlStore(db)
		business := bizemail.NewUpdateEmailBiz(store)

		if err := business.UpdateEmail(c.Request.Context(), id, &updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
