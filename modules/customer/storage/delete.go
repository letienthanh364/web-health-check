package storagecustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

func (s *sqlStore) DeleteCustomer(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "deleted"

	if err := s.db.Where(cond).Updates(&modelcustomer.CustomerUpdate{Status: &deletedStatus}).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
