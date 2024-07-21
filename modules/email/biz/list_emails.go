package bizemail

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
)

type ListEmailStorage interface {
	ListEmail(
		ctx context.Context,
		filter *modelemail.Filter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelemail.Email, error)
}

type listEmailBiz struct {
	store ListEmailStorage
}

func NewListEmailBiz(store ListEmailStorage) *listEmailBiz {
	return &listEmailBiz{store: store}
}

func (biz *listEmailBiz) ListEmail(ctx context.Context,
	filter *modelemail.Filter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelemail.Email, error) {
	data, err := biz.store.ListEmail(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
