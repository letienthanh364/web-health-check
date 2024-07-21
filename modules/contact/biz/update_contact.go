package bizecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
)

type UpdateContactStorage interface {
	GetContact(ctx context.Context, cond map[string]interface{}) (*modelcontact.Contact, error)
	UpdateContact(ctx context.Context, cond map[string]interface{}, updateData *modelcontact.ContactUpdate) error
}

type updateContactBiz struct {
	store UpdateContactStorage
}

func NewUpdateContactBiz(store UpdateContactStorage) *updateContactBiz {
	return &updateContactBiz{store: store}
}

func (biz *updateContactBiz) UpdateContact(ctx context.Context, configId int, updateData *modelcontact.ContactUpdate) error {
	data, err := biz.store.GetContact(ctx, map[string]interface{}{"id": configId})
	if err != nil {
		return appCommon.ErrCannotGetEntity(modelcontact.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelcontact.ErrContactIsDeleted
	}

	if err := updateData.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateContact(ctx, map[string]interface{}{"id": configId}, updateData); err != nil {
		return appCommon.ErrCannotDeleteEntity(modelcontact.EntityName, err)
	}

	return nil
}
