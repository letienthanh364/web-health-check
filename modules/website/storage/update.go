package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) UpdateWebsite(ctx context.Context, cond map[string]interface{}, updateData *modelwebsite.WebsiteUpdate) error {
	if err := s.db.Where(cond).Updates(&updateData).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
