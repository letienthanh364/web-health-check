package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	"strconv"
)

type ListContactsForWebsiteStorage interface {
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
}

type ListContactsStorage interface {
	ListContacts(
		ctx context.Context,
		filter *modelcontact.Filter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelcontact.Contact, error)
}

type listContactsForWebsiteBiz struct {
	websiteStorage ListContactsForWebsiteStorage
	contactStorage ListContactsStorage
}

func NewListContactsForWebsiteBiz(websiteStorage ListContactsForWebsiteStorage, contactStorage ListContactsStorage) *listContactsForWebsiteBiz {
	return &listContactsForWebsiteBiz{
		websiteStorage: websiteStorage,
		contactStorage: contactStorage,
	}
}

func (biz *listContactsForWebsiteBiz) ListContactsForWebsite(ctx context.Context, websiteId int,
	filter *modelcontact.Filter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelwebsite.WebsiteContact, error) {
	data, err := biz.websiteStorage.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return nil, appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return nil, modelwebsite.ErrWebsiteIsDeleted
	}

	filter.WebsiteId = strconv.Itoa(websiteId)
	paging.Process()

	contactList, err := biz.contactStorage.ListContacts(ctx, filter, paging)
	if err != nil {
		return nil, appCommon.ErrCannotListEntity(modelcontact.EntityName, err)
	}

	var res []modelwebsite.WebsiteContact
	for _, c := range contactList {
		contact := modelwebsite.WebsiteContact{
			Id:             c.Id,
			ContactAddress: c.Address,
			ContactMethod:  c.ContactMethod,
		}
		res = append(res, contact)
	}

	return res, nil
}
