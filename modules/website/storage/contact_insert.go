package storagecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelemail "github.com/teddlethal/web-health-check/modules/contact/model"
)

func (s *sqlStore) CreateContact(ctx context.Context, data *modelemail.ContactCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return appCommon.ErrDB(err)
	}

	return nil
}
