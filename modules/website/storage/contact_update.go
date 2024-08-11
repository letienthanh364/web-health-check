package storagecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/contact/model"
)

func (s *sqlStore) UpdateContact(ctx context.Context, cond map[string]interface{}, updateData *modelcontact.ContactUpdate) error {
	if err := s.db.Where(cond).Updates(&updateData).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
