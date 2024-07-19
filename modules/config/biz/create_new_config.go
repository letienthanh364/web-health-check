package bizconfig

import (
	"context"
	configmodel "github.com/teddlethal/web-health-check/modules/config/model"
)

type CreateConfigStorage interface {
	CreateConfig(ctx context.Context, data *configmodel.ConfigCreation) error
}

type createConfigBiz struct {
	store CreateConfigStorage
}

func NewCreateConfigBiz(store CreateConfigStorage) *createConfigBiz {
	return &createConfigBiz{store: store}
}

func (biz *createConfigBiz) CreateNewConfig(ctx context.Context, data *configmodel.ConfigCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateConfig(ctx, data); err != nil {
		return err
	}

	return nil
}
