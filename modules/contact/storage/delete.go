package storagecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/contact/model"
)

func (s *sqlStore) DeleteContact(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "deleted"

	if err := s.db.Where(cond).Updates(&modelcontact.ContactUpdate{Status: &deletedStatus}).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
