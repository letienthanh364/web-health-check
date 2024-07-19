package bizwebsite

import (
	"context"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type UpdateWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	UpdateWebsite(ctx context.Context, cond map[string]interface{}, updateData *modelwebsite.WebsiteUpdate) error
}

type updateWebsiteBiz struct {
	store UpdateWebsiteStorage
}

func NewUpdateWebsiteBiz(store UpdateWebsiteStorage) *updateWebsiteBiz {
	return &updateWebsiteBiz{store: store}
}

func (biz *updateWebsiteBiz) UpdateWebsite(ctx context.Context, configId int, updateData *modelwebsite.WebsiteUpdate) error {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": configId})
	if err != nil {
		return err
	}

	if data.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	if err := updateData.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateWebsite(ctx, map[string]interface{}{"id": configId}, updateData); err != nil {
		return err
	}

	return nil
}
