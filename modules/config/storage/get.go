package storageconfig

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) GetConfig(ctx context.Context, cond map[string]interface{}) (*modelconfig.Config, error) {
	var data modelconfig.Config

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
