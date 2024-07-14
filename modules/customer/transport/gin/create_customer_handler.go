package gincustomer

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizcustomer "github.com/teddlethal/web-health-check/modules/customer/biz"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
	storagecustomer "github.com/teddlethal/web-health-check/modules/customer/storage"
	"gorm.io/gorm"
	"net/http"
)

func CreateCustomer(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var createData modelcustomer.CustomerCreate

		if err := c.ShouldBind(&createData); err != nil {
			c.JSON(http.StatusBadRequest, appCommon.ErrInvalidRequest(err))
			return
		}

		store := storagecustomer.NewSqlStore(db)
		business := bizcustomer.NewCreateCustomerBiz(store)

		if err := business.CreateNewCustomer(c.Request.Context(), &createData); err != nil {
			c.JSON(http.StatusBadRequest, appCommon.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse(createData.Id))
	}
}
