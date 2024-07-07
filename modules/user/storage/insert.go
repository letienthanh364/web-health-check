package storage

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return appCommon.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return appCommon.ErrDB(err)
	}

	return nil
}
