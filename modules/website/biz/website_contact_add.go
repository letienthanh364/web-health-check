package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type CreateContactForWebsiteStorage interface {
	CreateWebsiteContact(ctx context.Context, data *modelwebsite.WebsiteContactCreation) error
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
}

type addContactForWebsiteBiz struct {
	store CreateContactForWebsiteStorage
}

func NewAddContactForWebsiteBiz(store CreateContactForWebsiteStorage) *addContactForWebsiteBiz {
	return &addContactForWebsiteBiz{
		store: store,
	}
}

func (biz *addContactForWebsiteBiz) AddContactForWebsite(ctx context.Context, websiteId int, contact *modelwebsite.WebsiteContactCreation) error {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	contactCreate := modelwebsite.WebsiteContactCreation{
		WebsiteId:     websiteId,
		Address:       contact.Address,
		ContactMethod: contact.ContactMethod,
	}
	if err := biz.store.CreateWebsiteContact(ctx, &contactCreate); err != nil {
		return appCommon.ErrCannotCreateEntity(modelwebsite.WebsiteContactEntity, err)
	}

	return nil
}
