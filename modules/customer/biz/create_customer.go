package bizcustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

type CreateCustomerStorage interface {
	CreateCustomer(ctx context.Context, data *modelcustomer.CustomerCreate) error
}

type createCustomerBiz struct {
	store CreateCustomerStorage
}

func NewCreateCustomerBiz(store CreateCustomerStorage) *createCustomerBiz {
	return &createCustomerBiz{store: store}
}

func (biz *createCustomerBiz) CreateNewCustomer(ctx context.Context, data *modelcustomer.CustomerCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateCustomer(ctx, data); err != nil {
		return appCommon.ErrCannotCreateEntity(modelcustomer.EntityName, err)
	}

	return nil
}
