package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	"strconv"
)

type ListCheckTimesForWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	ListCheckTimes(
		ctx context.Context,
		filter *modelwebsite.WebsiteCheckTimeFilter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelwebsite.WebsiteCheckTime, error)
}

type listCheckTimesForWebsiteBiz struct {
	store ListCheckTimesForWebsiteStorage
}

func NewListCheckTimesForWebsiteBiz(store ListCheckTimesForWebsiteStorage) *listCheckTimesForWebsiteBiz {
	return &listCheckTimesForWebsiteBiz{
		store: store,
	}
}

func (biz *listCheckTimesForWebsiteBiz) ListCheckTimesForWebsite(ctx context.Context, websiteId int,
	filter *modelwebsite.WebsiteCheckTimeFilter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelwebsite.WebsiteCheckTime, error) {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return nil, appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return nil, modelwebsite.ErrWebsiteIsDeleted
	}

	filter.WebsiteId = strconv.Itoa(websiteId)
	paging.Process()

	res, err := biz.store.ListCheckTimes(ctx, filter, paging)
	if err != nil {
		return nil, appCommon.ErrCannotListEntity("website check time", err)
	}

	return res, nil
}
