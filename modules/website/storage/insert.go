package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) CreateWebsite(ctx context.Context, data *modelwebsite.WebsiteCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
