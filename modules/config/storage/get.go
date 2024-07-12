package storage

import (
	"context"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) GetConfig(ctx context.Context, cond map[string]interface{}) (*modelconfig.Config, error) {
	var data modelconfig.Config

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
