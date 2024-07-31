package bizecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
)

type DeleteContactStorage interface {
	GetContact(ctx context.Context, cond map[string]interface{}) (*modelcontact.Contact, error)
	DeleteContact(ctx context.Context, cond map[string]interface{}) error
}

type deleteContactBiz struct {
	store DeleteContactStorage
}

func NewDeleteContactBiz(store DeleteContactStorage) *deleteContactBiz {
	return &deleteContactBiz{store: store}
}

func (biz *deleteContactBiz) DeleteEmailById(ctx context.Context, id int) error {
	data, err := biz.store.GetContact(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return appCommon.ErrCannotGetEntity(modelcontact.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelcontact.ErrContactIsDeleted
	}

	if err := biz.store.DeleteContact(ctx, map[string]interface{}{"id": id}); err != nil {
		return appCommon.ErrCannotDeleteEntity(modelcontact.EntityName, err)
	}

	return nil
}
