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

func ListContact(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelcontact.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storagecontact.NewSqlStore(db)
		business := bizecontact.NewListContactBiz(store)

		res, err := business.ListContacts(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, queryString.Filter, queryString.Paging))

	}
}
