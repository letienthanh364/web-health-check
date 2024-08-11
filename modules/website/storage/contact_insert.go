package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) CreateWebsiteContact(ctx context.Context, data *modelwebsite.WebsiteContactCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
