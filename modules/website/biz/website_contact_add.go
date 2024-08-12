package bizwebsite

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
	"strconv"
)

type CreateContactForWebsiteStorage interface {
	CreateWebsiteContact(ctx context.Context, data *modelwebsite.WebsiteContactCreation) error
	GetWebsite(ctx context.Context, cond map[string]interface{}) (*modelwebsite.Website, error)
	ListWebsiteContacts(
		ctx context.Context,
		filter *modelwebsite.WebsiteContactFilter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelwebsite.WebsiteContact, error)
}

type addContactForWebsiteBiz struct {
	store CreateContactForWebsiteStorage
}

func NewAddContactForWebsiteBiz(store CreateContactForWebsiteStorage) *addContactForWebsiteBiz {
	return &addContactForWebsiteBiz{
		store: store,
	}
}

func (biz *addContactForWebsiteBiz) AddContactForWebsite(ctx context.Context, websiteId int, contact *modelwebsite.WebsiteContactCreation) error {
	data, err := biz.store.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelwebsite.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelwebsite.ErrWebsiteIsDeleted
	}

	// Check if duplicate the default email
	if contact.Address == data.DefaultEmail {
		return modelwebsite.ErrContactExisted
	}

	// Check the contact store
	contactFilter := modelwebsite.WebsiteContactFilter{WebsiteId: strconv.Itoa(websiteId)}
	contactPaging := appCommon.Paging{Page: 1, Limit: modelwebsite.ContactLimit}

	contactList, err := biz.store.ListWebsiteContacts(ctx, &contactFilter, &contactPaging)

	if err != nil {
		return appCommon.ErrCannotListEntity(modelwebsite.WebsiteContactEntity, err)
	}

	if contactPaging.Total == modelwebsite.ContactLimit {
		return modelwebsite.ErrContactExceedLimit
	}

	for _, c := range contactList {
		if c.Address == contact.Address {
			return modelwebsite.ErrContactExisted
		}
	}

	// Add new contact
	contactCreate := modelwebsite.WebsiteContactCreation{
		WebsiteId:     websiteId,
		Address:       contact.Address,
		ContactMethod: contact.ContactMethod,
	}
	if err := biz.store.CreateWebsiteContact(ctx, &contactCreate); err != nil {
		return appCommon.ErrCannotCreateEntity(modelwebsite.WebsiteContactEntity, err)
	}

	return nil
}
