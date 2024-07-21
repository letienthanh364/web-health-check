package storageemail

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/email/model"
)

func (s *sqlStore) ListEmail(ctx context.Context,
	filter *modelemail.Filter,
	paging *appCommon.Paging,
	moreKeys ...string) ([]modelemail.Email, error) {
	var res []modelemail.Email

	db := s.db.Where("status <> ?", "deleted")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
		if v := f.WebsiteId; v != "" {
			db = db.Where("website_id = ?", v)
		}
	}

	if err := db.
		Table(modelemail.Email{}.TableName()).
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
