package storageconfig

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) ListConfig(ctx context.Context,
	filter *modelconfig.Filter,
	paging *appCommon.Paging,
	moreKeys ...string) ([]modelconfig.Config, error) {
	var res []modelconfig.Config

	db := s.db.Where("status <> ?", "deleted")

	//if filter != nil {
	//	if filter.Time != nil {
	//		log.Println("here")
	//		db = db.Where("time = ?", *filter.Time)
	//	}
	//	if filter.CheckLimit != nil {
	//		db = db.Where("check_limit = ?", *filter.CheckLimit)
	//	}
	//	if filter.StartTime != nil {
	//		db = db.Where("start_time = ?", *filter.StartTime)
	//	}
	//}

	if err := db.
		Table(modelconfig.Config{}.TableName()).
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
