package bizemail

import (
	"context"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
)

type UpdateEmailStorage interface {
	GetEmail(ctx context.Context, cond map[string]interface{}) (*modelemail.Email, error)
	UpdateEmail(ctx context.Context, cond map[string]interface{}, updateData *modelemail.EmailUpdate) error
}

type updateEmailBiz struct {
	store UpdateEmailStorage
}

func NewUpdateEmailBiz(store UpdateEmailStorage) *updateEmailBiz {
	return &updateEmailBiz{store: store}
}

func (biz *updateEmailBiz) UpdateEmail(ctx context.Context, configId int, updateData *modelemail.EmailUpdate) error {
	data, err := biz.store.GetEmail(ctx, map[string]interface{}{"id": configId})
	if err != nil {
		return err
	}

	if data.Status == "deleted" {
		return modelemail.ErrEmailIsDeleted
	}

	if err := updateData.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateEmail(ctx, map[string]interface{}{"id": configId}, updateData); err != nil {
		return err
	}

	return nil
}
