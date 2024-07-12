package storagecustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) DeleteCustomer(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "deleted"

	if err := s.db.Where(cond).Updates(&modelconfig.ConfigUpdate{Status: &deletedStatus}).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
