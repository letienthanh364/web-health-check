package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
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
	return &updateWebsiteBiz{
		store: store,
	}
}

func (biz *updateWebsiteBiz) UpdateWebsite(ctx context.Context, websiteId int, updateData *modelwebsite.WebsiteUpdate) error {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	// Check if website path is existed
	website, _ := biz.store.GetWebsite(ctx, map[string]interface{}{"path": updateData.Path})
	if website != nil && website.Id != websiteId {
		return modelwebsite.ErrPathIsExisted
	}

	// Update website
	if err := updateData.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateWebsite(ctx, map[string]interface{}{"id": websiteId}, updateData); err != nil {
		return appCommon.ErrCannotUpdateEntity(modelwebsite.EntityName, err)
	}

	return nil
}
