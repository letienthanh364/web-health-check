package bizecontact

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcontact "github.com/teddlethal/web-health-check/modules/contact/model"
)

type ListContactStorage interface {
	ListContacts(
		ctx context.Context,
		filter *modelcontact.Filter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelcontact.Contact, error)
}

type listContactBiz struct {
	store ListContactStorage
}

func NewListContactBiz(store ListContactStorage) *listContactBiz {
	return &listContactBiz{store: store}
}

func (biz *listContactBiz) ListContacts(ctx context.Context,
	filter *modelcontact.Filter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelcontact.Contact, error) {
	data, err := biz.store.ListContacts(ctx, filter, paging)

	if err != nil {
		return nil, appCommon.ErrCannotListEntity(modelcontact.EntityName, err)
	}

	return data, nil
}
