package storagecustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

func (s *sqlStore) GetCustomer(ctx context.Context, cond map[string]interface{}) (*modelcustomer.Customer, error) {
	var data modelcustomer.Customer

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
