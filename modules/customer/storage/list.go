package storagecustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

func (s *sqlStore) ListCustomer(ctx context.Context,
	filter *modelcustomer.Filter,
	paging *appCommon.Paging,
	moreKeys ...string) ([]modelcustomer.Customer, error) {
	var res []modelcustomer.Customer

	db := s.db.Where("status <> ?", "deleted")

	if err := db.
		Table(modelcustomer.Customer{}.TableName()).
		Select("id").
		Count(&paging.Total).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	if err := db.
		Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&res).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return res, nil
}
