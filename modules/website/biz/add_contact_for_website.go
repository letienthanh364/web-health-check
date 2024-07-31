package bizwebsite

import (
	"context"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
	modelwebsite "github.com/teddlethal/web-health-check/modules/website/model"
)

type CreateContactStorage interface {
	CreateContact(ctx context.Context, data *modelcontact.ContactCreation) error
}

type addContactForWebsiteBiz struct {
	contactStorage CreateContactStorage
	websiteStorage UpdateWebsiteStorage
}

func NewAddContactForWebsiteBiz(contactStorage CreateContactStorage, websiteStorage UpdateWebsiteStorage) *addContactForWebsiteBiz {
	return &addContactForWebsiteBiz{
		contactStorage: contactStorage,
		websiteStorage: websiteStorage,
	}
}

func (biz *addContactForWebsiteBiz) AddContactForWebsite(ctx context.Context, websiteId int, contact *modelwebsite.WebsiteContactCreation) (int, error) {
	returnId := -1
	data, err := biz.websiteStorage.GetWebsite(ctx, map[string]interface{}{"id": websiteId})
	if err != nil {
		return returnId, err
	}

	if data.Status == "deleted" {
		return returnId, modelwebsite.ErrWebsiteIsDeleted
	}

	contactCreate := modelcontact.ContactCreation{
		WebsiteId:     websiteId,
		Address:       contact.ContactAddress,
		ContactMethod: contact.ContactMethod,
	}
	if err := biz.contactStorage.CreateContact(ctx, &contactCreate); err != nil {
		return returnId, err
	}
	returnId = contactCreate.Id

	return returnId, nil
}
