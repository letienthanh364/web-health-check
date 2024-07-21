package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var createData modelwebsite.WebsiteCreation

		if err := c.ShouldBind(&createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewCreateWebsiteBiz(store)

		if err := business.CreateNewWebsite(c.Request.Context(), &createData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, appCommon.SimpleSuccessResponse(createData.Id))
	}
}
