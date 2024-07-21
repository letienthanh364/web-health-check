package bizecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
)

type CreateContactStorage interface {
	CreateContact(ctx context.Context, data *modelcontact.ContactCreation) error
}

type createContactBiz struct {
	store CreateContactStorage
}

func NewCreateContactBiz(store CreateContactStorage) *createContactBiz {
	return &createContactBiz{store: store}
}

func (biz *createContactBiz) CreateNewContact(ctx context.Context, data *modelcontact.ContactCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateContact(ctx, data); err != nil {
		return appCommon.ErrCannotCreateEntity(modelcontact.EntityName, err)
	}

	return nil
}
