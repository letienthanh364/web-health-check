package storage

import (
	"context"
	configmodel "github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) CreateConfig(ctx context.Context, data *configmodel.ConfigCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
