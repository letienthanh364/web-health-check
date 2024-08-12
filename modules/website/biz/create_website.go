package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type CreateWebsiteStorage interface {
	CreateWebsite(ctx context.Context, data *modelwebsite.WebsiteCreation) error
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
}

type createWebsiteBiz struct {
	store CreateWebsiteStorage
}

func NewCreateWebsiteBiz(store CreateWebsiteStorage) *createWebsiteBiz {
	return &createWebsiteBiz{store: store}
}

func (biz *createWebsiteBiz) CreateNewWebsite(ctx context.Context, data *modelwebsite.WebsiteCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// Check if website path is existed
	website, _ := biz.store.GetWebsite(ctx, map[string]interface{}{"path": data.Path})
	if website != nil {
		return modelwebsite.ErrPathExisted
	}

	// Create new website
	if err := biz.store.CreateWebsite(ctx, data); err != nil {
		return appCommon.ErrCannotCreateEntity(modelwebsite.EntityName, err)
	}

	return nil
}
