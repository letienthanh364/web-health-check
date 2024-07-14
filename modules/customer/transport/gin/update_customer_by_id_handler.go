package gincustomer

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizcustomer "github.com/teddlethal/web-health-check/modules/customer/biz"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
	storagecustomer "github.com/teddlethal/web-health-check/modules/customer/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateCustomer(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var updateData modelcustomer.CustomerUpdate

		if err := c.ShouldBind(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, appCommon.ErrInvalidRequest(err))
			return
		}

		requester := c.MustGet(appCommon.CurrentUser).(appCommon.Requester)
		store := storagecustomer.NewSqlStore(db)
		business := bizcustomer.NewUpdateCustomerBiz(store, requester)

		if err := business.UpdateCustomerById(c.Request.Context(), id, &updateData); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
