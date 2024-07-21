package storageemail

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/email/model"
)

func (s *sqlStore) DeleteEmail(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "deleted"

	if err := s.db.Where(cond).Updates(&modelemail.EmailUpdate{Status: &deletedStatus}).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
