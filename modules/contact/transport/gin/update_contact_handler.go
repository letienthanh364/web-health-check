package gincontact

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizemail "github.com/teddlethal/web-health-check/modules/contact/biz"
	modelemail "github.com/teddlethal/web-health-check/modules/contact/model"
	storageemail "github.com/teddlethal/web-health-check/modules/contact/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var updateData modelemail.ContactUpdate

		if err := c.ShouldBind(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storageemail.NewSqlStore(db)
		business := bizemail.NewUpdateContactBiz(store)

		if err := business.UpdateContact(c.Request.Context(), id, &updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
