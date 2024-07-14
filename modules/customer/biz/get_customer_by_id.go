package bizcustomer

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	modelcustomer "github.com/teddlethal/web-health-check/modules/customer/model"
)

type GetCustomerStorage interface {
	GetCustomer(ctx context.Context, cond map[string]interface{}) (*modelcustomer.Customer, error)
}

type getCustomerBiz struct {
	store GetCustomerStorage
}

func NewGetCustomerBiz(store GetCustomerStorage) *getCustomerBiz {
	return &getCustomerBiz{store: store}
}

func (biz *getCustomerBiz) GetCustomerById(ctx context.Context, id int) (*modelcustomer.Customer, error) {
	data, err := biz.store.GetCustomer(ctx, map[string]interface{}{"id": id})

	if data.Status == "deleted" {
		return nil, modelcustomer.ErrCustomerIsDeleted
	}

	if err != nil {
		return nil, appCommon.ErrCannotGetEntity(modelcustomer.EntityName, err)
	}
	return data, nil
}
