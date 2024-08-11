package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) GetWebsiteContact(ctx context.Context, cond map[string]interface{}) (*modelwebsite.WebsiteContact, error) {
	var data modelwebsite.WebsiteContact

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
