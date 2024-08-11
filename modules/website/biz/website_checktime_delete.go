package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type DeleteCheckTimeForWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	DeleteWebsiteCheckTime(ctx context.Context, cond map[string]interface{}) error
}

type deleteCheckTimeForWebsiteBiz struct {
	store DeleteCheckTimeForWebsiteStorage
}

func NewDeleteCheckTimeForWebsiteBiz(store DeleteCheckTimeForWebsiteStorage) *deleteCheckTimeForWebsiteBiz {
	return &deleteCheckTimeForWebsiteBiz{
		store: store,
	}
}

func (biz *deleteCheckTimeForWebsiteBiz) DeleteCheckTimeForWebsite(ctx context.Context, websiteId int, checkTimeId int) error {
	websiteData, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if websiteData.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	if err := biz.store.DeleteWebsiteCheckTime(ctx, map[string]interface{}{"id": checkTimeId}); err != nil {
		return appCommon.ErrCannotDeleteEntity("website check time", err)
	}

	return nil
}
