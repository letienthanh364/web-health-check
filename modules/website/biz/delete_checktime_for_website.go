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

func (biz *deleteCheckTimeForWebsiteBiz) DeleteCheckTimeForWebsite(ctx context.Context, checktimeId int) error {
	if err := biz.store.DeleteWebsiteCheckTime(ctx, map[string]interface{}{"id": checktimeId}); err != nil {
		return appCommon.ErrCannotDeleteEntity("website check time", err)
	}

	return nil
}
