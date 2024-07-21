package storageemail

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/email/model"
)

func (s *sqlStore) GetEmail(ctx context.Context, cond map[string]interface{}) (*modelemail.Email, error) {
	var data modelemail.Email

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
