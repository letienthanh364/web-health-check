package bizcustomer

import (
	"context"
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

type UpdateCustomerStore interface {
	GetCustomer(ctx context.Context, cond map[string]interface{}) (*modelcustomer.Customer, error)
	UpdateCustomer(ctx context.Context, cond map[string]interface{}, updateData *modelcustomer.CustomerUpdate) error
}

type updateCustomerBiz struct {
	store     UpdateCustomerStore
	requester appCommon.Requester
}

func NewUpdateCustomerBiz(store UpdateCustomerStore, requester appCommon.Requester) *updateCustomerBiz {
	return &updateCustomerBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *updateCustomerBiz) UpdateCustomerById(ctx context.Context, id int, updateData *modelcustomer.CustomerUpdate) error {
	data, err := biz.store.GetCustomer(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return appCommon.ErrCannotGetEntity(modelcustomer.EntityName, err)
	}

	if data.Status == "deleted" {
		return modelcustomer.ErrCustomerIsDeleted
	}

	if !appCommon.IsAdmin(biz.requester) {
		return appCommon.ErrNoPermission(errors.New("no permission"))
	}

	if err := biz.store.UpdateCustomer(ctx, map[string]interface{}{"id": id}, updateData); err != nil {
		return appCommon.ErrCannotDeleteEntity(modelcustomer.EntityName, err)
	}

	return nil
}
