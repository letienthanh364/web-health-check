package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type CreateCheckTimeForWebsiteStorage interface {
	CreateWebsiteCheckTime(ctx context.Context, data *modelwebsite.WebsiteCheckTimeCreation) error
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
}

type addWebsiteCheckTimeForWebsiteBiz struct {
	store CreateCheckTimeForWebsiteStorage
}

func NewAddCheckTimeForWebsiteBiz(store CreateCheckTimeForWebsiteStorage) *addWebsiteCheckTimeForWebsiteBiz {
	return &addWebsiteCheckTimeForWebsiteBiz{
		store: store,
	}
}

func (biz *addWebsiteCheckTimeForWebsiteBiz) AddCheckTimeForWebsite(ctx context.Context, websiteId int, data *modelwebsite.WebsiteCheckTimeCreation) error {
	website, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if website.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	contactCreate := modelwebsite.WebsiteCheckTimeCreation{
		WebsiteId: websiteId,
		CheckTime: data.CheckTime,
	}
	if err := biz.store.CreateWebsiteCheckTime(ctx, &contactCreate); err != nil {
		return err
	}

	return nil
}
