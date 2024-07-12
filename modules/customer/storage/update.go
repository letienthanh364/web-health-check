package storagecustomer

import (
	"context"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

func (s *sqlStore) UpdateCustomer(ctx context.Context, cond map[string]interface{}, updateData *modelcustomer.CustomerUpdate) error {
	if err := s.db.Where(cond).Updates(&updateData).Error; err != nil {
		return err
	}

	return nil
}
