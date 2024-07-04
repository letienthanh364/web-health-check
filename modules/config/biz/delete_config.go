package biz

import (
	"context"
	configmodel "github.com/teddlethal/web-health-check/modules/config/model"
)

type DeleteConfigStorage interface {
	GetConfig(ctx context.Context, cond map[string]interface{}) (*configmodel.Config, error)
	DeleteConfig(ctx context.Context, cond map[string]interface{}) error
}

type deleteConfigBiz struct {
	store DeleteConfigStorage
}

func NewDeleteConfigBiz(store DeleteConfigStorage) *deleteConfigBiz {
	return &deleteConfigBiz{store: store}
}

func (biz *deleteConfigBiz) DeleteConfigById(ctx context.Context, configId int) error {
	data, err := biz.store.GetConfig(ctx, map[string]interface{}{"id": configId})

	if err != nil {
		return err
	}

	if data.Status == "deleted" {
		return configmodel.ErrConfigIsDeleted
	}

	if err := biz.store.DeleteConfig(ctx, map[string]interface{}{"id": configId}); err != nil {
		return err
	}

	return nil
}
