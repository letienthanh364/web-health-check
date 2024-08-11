package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) DeleteWebsiteCheckTime(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(modelwebsite.WebsiteCheckTime{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
