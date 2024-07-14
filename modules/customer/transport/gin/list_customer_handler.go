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

func ListCustomer(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var queryString struct {
			appCommon.Paging
			modelcustomer.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		requester := c.MustGet(appCommon.CurrentUser).(appCommon.Requester)
		store := storagecustomer.NewSqlStore(db)
		business := bizcustomer.NewListCustomerBiz(store, requester)

		res, err := business.ListCustomer(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, appCommon.NewSuccessResponse(res, queryString.Filter, queryString.Paging))

	}
}
