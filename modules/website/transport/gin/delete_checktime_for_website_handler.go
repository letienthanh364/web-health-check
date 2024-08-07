package ginwebsite

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizwebsite "github.com/teddlethal/web-health-check/modules/website/biz"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	storagewebsite "github.com/teddlethal/web-health-check/modules/website/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteCheckTimeForWebsite(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		webId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var deleteData modelwebsite.WebsiteCheckTimeDelete

		if err := c.ShouldBind(&deleteData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storagewebsite.NewSqlStore(db)
		business := bizwebsite.NewDeleteCheckTimeForWebsiteBiz(store)

		if err := business.DeleteCheckTimeForWebsite(c.Request.Context(), webId, deleteData.Id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
