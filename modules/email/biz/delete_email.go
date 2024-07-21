package bizemail

import (
	"context"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
)

type DeleteEmailStorage interface {
	GetEmail(ctx context.Context, cond map[string]interface{}) (*modelemail.Email, error)
	DeleteEmail(ctx context.Context, cond map[string]interface{}) error
}

type deleteEmailBiz struct {
	store DeleteEmailStorage
}

func NewDeleteEmailBiz(store DeleteEmailStorage) *deleteEmailBiz {
	return &deleteEmailBiz{store: store}
}

func (biz *deleteEmailBiz) DeleteEmailById(ctx context.Context, configId int) error {
	data, err := biz.store.GetEmail(ctx, map[string]interface{}{"id": configId})

	if err != nil {
		return err
	}

	if data.Status == "deleted" {
		return modelemail.ErrEmailIsDeleted
	}

	if err := biz.store.DeleteEmail(ctx, map[string]interface{}{"id": configId}); err != nil {
		return err
	}

	return nil
}
