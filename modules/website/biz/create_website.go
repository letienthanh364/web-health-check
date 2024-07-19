package bizwebsite

import (
	"context"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type CreateWebsiteStorage interface {
	CreateWebsite(ctx context.Context, data *modelwebsite.WebsiteCreation) error
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

	if err := biz.store.CreateWebsite(ctx, data); err != nil {
		return err
	}

	return nil
}
