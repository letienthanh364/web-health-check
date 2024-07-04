package storage

import (
	"context"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

func (s *sqlStore) DeleteConfig(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := "Deleted"

	if err := s.db.Where(cond).Updates(&configmodel.ConfigUpdate{Status: &deletedStatus}).Error; err != nil {
		return err
	}

	return nil
}
