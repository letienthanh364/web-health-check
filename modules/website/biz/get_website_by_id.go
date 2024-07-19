package bizwebsite

import (
	"context"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type GetWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
}

type getWebsiteBiz struct {
	store GetWebsiteStorage
}

func NewGetWebsiteBiz(store GetWebsiteStorage) *getWebsiteBiz {
	return &getWebsiteBiz{store: store}
}

func (biz *getWebsiteBiz) GetWebsiteById(ctx context.Context, id int) (*modelwebsite.Website, error) {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if data.Status == "deleted" {
		return nil, modelwebsite.ErrWebsiteIsDeleted
	}

	return data, nil
}
