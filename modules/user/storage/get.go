package storageuser

import (
	"context"
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	modeluser "github.com/teddlethal/web-health-check/modules/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, cond map[string]interface{}, moreInfo ...string) (*modeluser.User, error) {
	db := s.db.Table(modeluser.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user modeluser.User

	if err := s.db.Where(cond).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appCommon.ErrRecordNotFound
		}

		return nil, appCommon.ErrDB(err)
	}

	return &user, nil
}
