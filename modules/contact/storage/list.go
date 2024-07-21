package storagecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
)

func (s *sqlStore) ListContacts(ctx context.Context,
	filter *modelcontact.Filter,
	paging *appCommon.Paging,
	moreKeys ...string) ([]modelcontact.Contact, error) {
	var res []modelcontact.Contact

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
		Table(modelcontact.Contact{}.TableName()).
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
