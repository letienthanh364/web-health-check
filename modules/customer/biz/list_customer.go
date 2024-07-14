package bizcustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

type ListCustomerStorage interface {
	ListCustomer(
		ctx context.Context,
		filter *modelcustomer.Filter,
		paging *appCommon.Paging,
		moreKeys ...string,
	) ([]modelcustomer.Customer, error)
}

type listCustomerBiz struct {
	store     ListCustomerStorage
	requester appCommon.Requester
}

func NewListCustomerBiz(store ListCustomerStorage, requester appCommon.Requester) *listCustomerBiz {
	return &listCustomerBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *listCustomerBiz) ListCustomer(ctx context.Context,
	filter *modelcustomer.Filter,
	paging *appCommon.Paging,
	moreKeys ...string,
) ([]modelcustomer.Customer, error) {
	ctxStore := context.WithValue(ctx, appCommon.CurrentUser, biz.requester)

	data, err := biz.store.ListCustomer(ctxStore, filter, paging)

	if err != nil {
		return nil, appCommon.ErrCannotListEntity(modelcustomer.EntityName, err)
	}

	return data, nil
}
