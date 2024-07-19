package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type ListWebsiteStorage interface {
	ListWebsite(
		ctx context.Context,
		filter *modelwebsite.Filter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelwebsite.Website, error)
}

type listWebsiteBiz struct {
	store ListWebsiteStorage
}

func NewListWebsiteBiz(store ListWebsiteStorage) *listWebsiteBiz {
	return &listWebsiteBiz{store: store}
}

func (biz *listWebsiteBiz) ListWebsite(ctx context.Context,
	filter *modelwebsite.Filter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelwebsite.Website, error) {
	data, err := biz.store.ListWebsite(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
