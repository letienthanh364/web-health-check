package gincontact

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/contact/biz"
	"github.com/teddlethal/web-health-check/modules/contact/model"
	"github.com/teddlethal/web-health-check/modules/contact/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var createData modelcontact.ContactCreation

		if err := c.ShouldBind(&createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storagecontact.NewSqlStore(db)
		business := bizecontact.NewCreateContactBiz(store)

		if err := business.CreateNewContact(c.Request.Context(), &createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, appCommon.SimpleSuccessResponse(createData.Id))
	}
}
