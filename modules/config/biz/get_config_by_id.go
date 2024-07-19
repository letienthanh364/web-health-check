package bizconfig

import (
	"context"
	configmodel "github.com/teddlethal/web-health-check/modules/config/model"
)

type GetConfigStorage interface {
	GetConfig(ctx context.Context, cond map[string]interface{}) (*configmodel.Config, error)
}

type getConfigBiz struct {
	store GetConfigStorage
}

func NewGetConfigBiz(store GetConfigStorage) *getConfigBiz {
	return &getConfigBiz{store: store}
}

func (biz *getConfigBiz) GetConfigById(ctx context.Context, id int) (*configmodel.Config, error) {
	data, err := biz.store.GetConfig(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if data.Status == "deleted" {
		return nil, configmodel.ErrConfigIsDeleted
	}

	return data, nil
}
