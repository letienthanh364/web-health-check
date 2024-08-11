package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type DeleteContactForWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	DeleteWebsiteContact(ctx context.Context, cond map[string]interface{}) error
}

type deleteContactForWebsiteBiz struct {
	store DeleteContactForWebsiteStorage
}

func NewDeleteContactForWebsiteBiz(store DeleteContactForWebsiteStorage) *deleteContactForWebsiteBiz {
	return &deleteContactForWebsiteBiz{
		store: store,
	}
}

func (biz *deleteContactForWebsiteBiz) DeleteContactForWebsite(ctx context.Context, websiteId int, contactId int) error {
	websiteData, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if websiteData.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	if err := biz.store.DeleteWebsiteContact(ctx, map[string]interface{}{"id": contactId}); err != nil {
		return appCommon.ErrCannotDeleteEntity(modelwebsite.EntityName, err)
	}

	return nil
}
