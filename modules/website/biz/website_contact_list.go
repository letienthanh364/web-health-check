package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	"strconv"
)

type ListContactsForWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	ListWebsiteContacts(
		ctx context.Context,
		filter *modelwebsite.WebsiteContactFilter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelwebsite.WebsiteContact, error)
}

type listContactsForWebsiteBiz struct {
	store ListContactsForWebsiteStorage
}

func NewListContactsForWebsiteBiz(store ListContactsForWebsiteStorage) *listContactsForWebsiteBiz {
	return &listContactsForWebsiteBiz{
		store: store,
	}
}

func (biz *listContactsForWebsiteBiz) ListContactsForWebsite(ctx context.Context, websiteId int,
	filter *modelwebsite.WebsiteContactFilter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelwebsite.WebsiteContact, error) {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return nil, appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return nil, modelwebsite.ErrWebsiteIsDeleted
	}

	filter.WebsiteId = strconv.Itoa(websiteId)
	paging.Process()

	contactList, err := biz.store.ListWebsiteContacts(ctx, filter, paging)
	if err != nil {
		return nil, appCommon.ErrCannotListEntity(modelwebsite.EntityName, err)
	}

	return contactList, nil
}
