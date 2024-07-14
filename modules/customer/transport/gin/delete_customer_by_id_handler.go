package gincustomer

import (
	"github.com/gin-gonic/gin"
	"github.com/teddlethal/web-health-check/appCommon"
	bizcustomer "github.com/teddlethal/web-health-check/modules/customer/biz"
	storagecustomer "github.com/teddlethal/web-health-check/modules/customer/storage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteCustomer(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storagecustomer.NewSqlStore(db)
		business := bizcustomer.NewDeleteCustomerBiz(store)

		if err := business.DeleteCustomerById(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, appCommon.SimpleSuccessResponse("success"))
	}
}
