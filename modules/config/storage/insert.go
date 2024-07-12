package storageconfig

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) CreateConfig(ctx context.Context, data *modelconfig.ConfigCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
