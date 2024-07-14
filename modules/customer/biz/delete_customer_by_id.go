package bizcustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

type DeleteCustomerStore interface {
	GetCustomer(ctx context.Context, cond map[string]interface{}) (*modelcustomer.Customer, error)
	DeleteCustomer(ctx context.Context, cond map[string]interface{}) error
}

type deleteCustomerBiz struct {
	store DeleteCustomerStore
}

func NewDeleteCustomerBiz(store DeleteCustomerStore) *deleteCustomerBiz {
	return &deleteCustomerBiz{store: store}
}

func (biz *deleteCustomerBiz) DeleteCustomerById(ctx context.Context, id int) error {
	data, err := biz.store.GetCustomer(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return appCommon.ErrCannotGetEntity(modelcustomer.EntityName, err)
	}

	if data.Status == "deleted" {
		return appCommon.ErrEntityDeleted(modelcustomer.EntityName, modelcustomer.ErrCustomerIsDeleted)
	}

	if err := biz.store.DeleteCustomer(ctx, map[string]interface{}{"id": id}); err != nil {
		return appCommon.ErrCannotDeleteEntity(modelcustomer.EntityName, err)
	}

	return nil
}
