package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) CreateWebsiteCheckTime(ctx context.Context, data *modelwebsite.WebsiteCheckTimeCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
