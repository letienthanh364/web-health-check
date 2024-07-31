package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type DeleteContactStorage interface {
	GetContact(ctx context.Context, cond map[string]interface{}) (*modelcontact.Contact, error)
	DeleteContact(ctx context.Context, cond map[string]interface{}) error
}

type deleteContactForWebsiteBiz struct {
	websiteStorage GetWebsiteStorage
	contactStorage DeleteContactStorage
}

func NewDeleteContactForWebsiteBiz(websiteStorage GetWebsiteStorage, contactStorage DeleteContactStorage) *deleteContactForWebsiteBiz {
	return &deleteContactForWebsiteBiz{
		websiteStorage: websiteStorage,
		contactStorage: contactStorage,
	}
}

func (biz *deleteContactForWebsiteBiz) DeleteContactForWebsite(ctx context.Context, websiteId int, contactId int) error {
	websiteData, err := biz.websiteStorage.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if websiteData.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	contactData, err := biz.contactStorage.GetContact(ctx, map[string]interface{}{"id": contactId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelcontact.EntityName, err)
	}
	if contactData.Status == "deleted" {
		return modelcontact.ErrContactIsDeleted
	}

	return nil
}
