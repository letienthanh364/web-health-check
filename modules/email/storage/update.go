package storageemail

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/email/model"
)

func (s *sqlStore) UpdateEmail(ctx context.Context, cond map[string]interface{}, updateData *modelemail.EmailUpdate) error {
	if err := s.db.Where(cond).Updates(&updateData).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
