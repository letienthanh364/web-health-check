package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type UpdateWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	UpdateWebsite(ctx context.Context, cond map[string]interface{}, updateData *modelwebsite.WebsiteUpdate) error
	ListCheckTimes(
		ctx context.Context,
		filter *modelwebsite.WebsiteCheckTimeFilter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelwebsite.WebsiteCheckTime, error)
}

type updateWebsiteBiz struct {
	websiteStore UpdateWebsiteStorage
	contactStore ListContactsStorage
}

func NewUpdateWebsiteBiz(websiteStore UpdateWebsiteStorage, contactStore ListContactsStorage) *updateWebsiteBiz {
	return &updateWebsiteBiz{
		websiteStore: websiteStore,
		contactStore: contactStore,
	}
}

func (biz *updateWebsiteBiz) UpdateWebsite(ctx context.Context, websiteId int, updateData *modelwebsite.WebsiteUpdate) error {
	data, err := biz.websiteStore.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	if err := updateData.Validate(); err != nil {
		return err
	}

	if err := biz.websiteStore.UpdateWebsite(ctx, map[string]interface{}{"id": websiteId}, updateData); err != nil {
		return appCommon.ErrCannotUpdateEntity(modelwebsite.EntityName, err)
	}

	return nil
}
