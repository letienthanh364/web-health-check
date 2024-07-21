package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	modelemail "github.com/teddlethal/web-health-check/modules/contact/model"
	storageemail "github.com/teddlethal/web-health-check/modules/contact/storage"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func ListContactsForWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelemail.Filter
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		websiteStorage := storagewebsite.NewSqlStore(db)
		emailStorage := storageemail.NewSqlStore(db)
		business := bizwebsite.NewListContactsForWebsiteBiz(websiteStorage, emailStorage)

		res, err := business.ListContactsForWebsite(c.Request.Context(), id, &queryString.Filter, &queryString.Paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, appCommon.SimpleSuccessResponse(res))
	}
}
