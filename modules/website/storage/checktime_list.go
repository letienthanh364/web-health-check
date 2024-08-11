package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) ListCheckTimes(ctx context.Context,
	filter *modelwebsite.WebsiteCheckTimeFilter,
	paging *appCommon.Paging,
	moreKeys ...string) ([]modelwebsite.WebsiteCheckTime, error) {
	var res []modelwebsite.WebsiteCheckTime

	db := s.db

	if f := filter; f != nil {
		if v := f.WebsiteId; v != "" {
			db = db.Where("website_id = ?", v)
		}
	}

	if err := db.
		Table(modelwebsite.WebsiteCheckTime{}.TableName()).
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
