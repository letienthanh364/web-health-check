package bizconfig

import (
	"context"
	configmodel "github.com/teddlethal/web-health-check/modules/config/model"
)

type UpdateConfigStorage interface {
	GetConfig(ctx context.Context, cond map[string]interface{}) (*configmodel.Config, error)
	UpdateConfig(ctx context.Context, cond map[string]interface{}, updateData *configmodel.ConfigUpdate) error
}

type updateConfigBiz struct {
	store UpdateConfigStorage
}

func NewUpdateConfigBiz(store UpdateConfigStorage) *updateConfigBiz {
	return &updateConfigBiz{store: store}
}

func (biz *updateConfigBiz) UpdateConfig(ctx context.Context, configId int, updateData *configmodel.ConfigUpdate) error {
	data, err := biz.store.GetConfig(ctx, map[string]interface{}{"id": configId})
	if err != nil {
		return err
	}

	if data.Status == "deleted" {
		return configmodel.ErrConfigIsDeleted
	}

	if err := updateData.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateConfig(ctx, map[string]interface{}{"id": configId}, updateData); err != nil {
		return err
	}

	return nil
}
