package storagewebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

func (s *sqlStore) DeleteWebsite(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "deleted"

	if err := s.db.Where(cond).Updates(&modelwebsite.WebsiteUpdate{Status: &deletedStatus}).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
