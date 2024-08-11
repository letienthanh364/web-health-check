package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/website/biz"
	"github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AddContactForWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var addContactData modelwebsite.WebsiteContactCreation

		if err := c.ShouldBind(&addContactData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := addContactData.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		websiteStorage := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewAddContactForWebsiteBiz(websiteStorage)

		if err := business.AddContactForWebsite(c.Request.Context(), id, &addContactData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
