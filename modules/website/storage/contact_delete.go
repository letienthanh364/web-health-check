package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) DeleteWebsiteContact(ctx context.Context, cond map[string]interface{}) error {

	if err := s.db.Table(modelwebsite.WebsiteContact{}.TableName()).Where(cond).Delete(nil).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
