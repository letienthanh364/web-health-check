package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	"strconv"
)

type CreateCheckTimeForWebsiteStorage interface {
	CreateWebsiteCheckTime(ctx context.Context, data *modelwebsite.WebsiteCheckTimeCreation) error
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	ListCheckTimes(
		ctx context.Context,
		filter *modelwebsite.WebsiteCheckTimeFilter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelwebsite.WebsiteCheckTime, error)
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

	// Check the checktime store
	checktimeFilter := modelwebsite.WebsiteCheckTimeFilter{WebsiteId: strconv.Itoa(websiteId)}
	checktimePaging := appCommon.Paging{Page: 1, Limit: modelwebsite.CheckTimeLimit}

	checktimeList, err := biz.store.ListCheckTimes(ctx, &checktimeFilter, &checktimePaging)

	if err != nil {
		return appCommon.ErrCannotListEntity(modelwebsite.WebsiteCheckTimeEntity, err)
	}

	if checktimePaging.Total == modelwebsite.CheckTimeLimit {
		return modelwebsite.ErrCheckTimeExceedLimit
	}

	for _, c := range checktimeList {
		if c.CheckTime == data.CheckTime {
			return modelwebsite.ErrCheckTimeExisted
		}
	}

	// Add new check time
	checktimeCreate := modelwebsite.WebsiteCheckTimeCreation{
		WebsiteId: websiteId,
		CheckTime: data.CheckTime,
	}
	if err := biz.store.CreateWebsiteCheckTime(ctx, &checktimeCreate); err != nil {
		return appCommon.ErrCannotCreateEntity(modelwebsite.WebsiteCheckTimeEntity, err)
	}

	return nil
}
