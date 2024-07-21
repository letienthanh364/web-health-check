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

func ListEmail(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelemail.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		store := storageemail.NewSqlStore(db)
		business := bizemail.NewListEmailBiz(store)

		res, err := business.ListEmail(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, queryString.Filter, queryString.Paging))

	}
}
