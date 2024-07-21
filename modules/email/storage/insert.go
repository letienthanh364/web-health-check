package storageemail

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
)

func (s *sqlStore) CreateEmail(ctx context.Context, data *modelemail.EmailCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
