package storagecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/contact/model"
)

func (s *sqlStore) GetContact(ctx context.Context, cond map[string]interface{}) (*modelcontact.Contact, error) {
	var data modelcontact.Contact

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, appCommon.ErrDB(err)
	}

	return &data, nil
}
