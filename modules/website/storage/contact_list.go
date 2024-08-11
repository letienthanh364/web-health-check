package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) ListWebsiteContacts(ctx context.Context,
	filter *modelwebsite.WebsiteContactFilter,
	paging *appCommon.Paging,
	moreKeys ...string) ([]modelwebsite.WebsiteContact, error) {
	var res []modelwebsite.WebsiteContact

	db := s.db

	if f := filter; f != nil {

		if v := f.WebsiteId; v != "" {
			db = db.Where("website_id = ?", v)
		}
	}

	if err := db.
		Table(modelwebsite.WebsiteContact{}.TableName()).
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
