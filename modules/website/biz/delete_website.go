package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type DeleteWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	DeleteWebsite(ctx context.Context, cond map[string]interface{}) error
}

type deleteWebsiteBiz struct {
	store DeleteWebsiteStorage
}

func NewDeleteWebsiteBiz(store DeleteWebsiteStorage) *deleteWebsiteBiz {
	return &deleteWebsiteBiz{store: store}
}

func (biz *deleteWebsiteBiz) DeleteWebsiteById(ctx context.Context, configId int) error {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": configId})

	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	if err := biz.store.DeleteWebsite(ctx, map[string]interface{}{"id": configId}); err != nil {
		return appCommon.ErrCannotDeleteEntity(modelwebsite.EntityName, err)
	}

	return nil
}
