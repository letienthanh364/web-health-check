package bizecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
)

type GetContactStorage interface {
	GetContact(ctx context.Context, cond map[string]interface{}) (*modelcontact.Contact, error)
}

type getContactBiz struct {
	store GetContactStorage
}

func NewGetContactBiz(store GetContactStorage) *getContactBiz {
	return &getContactBiz{store: store}
}

func (biz *getContactBiz) GetContactById(ctx context.Context, id int) (*modelcontact.Contact, error) {
	data, err := biz.store.GetContact(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if data.Status == "deleted" {
		return nil, appCommon.ErrCannotDeleteEntity(modelcontact.EntityName, err)
	}

	return data, nil
}
