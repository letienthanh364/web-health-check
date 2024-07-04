package biz

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	"github.com/teddlethal/web-health-check/modules/config/model"
)

type ListConfigStorage interface {
	ListConfig(
		ctx context.Context,
		filter *configmodel.Filter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]configmodel.Config, error)
}

type listConfigBiz struct {
	store ListConfigStorage
}

func NewListConfigBiz(store ListConfigStorage) *listConfigBiz {
	return &listConfigBiz{store: store}
}

func (biz *listConfigBiz) ListConfig(ctx context.Context,
	filter *configmodel.Filter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]configmodel.Config, error) {
	data, err := biz.store.ListConfig(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
