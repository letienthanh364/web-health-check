package storage

import (
	"context"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) UpdateConfig(ctx context.Context, cond map[string]interface{}, updateData *modelconfig.ConfigUpdate) error {
	if err := s.db.Where(cond).Updates(&updateData).Error; err != nil {
		return err
	}

	return nil
}
