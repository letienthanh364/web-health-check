package bizemail

import (
	"context"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
)

type CreateEmailStorage interface {
	CreateEmail(ctx context.Context, data *modelemail.EmailCreation) error
}

type createEmailBiz struct {
	store CreateEmailStorage
}

func NewCreateEmailBiz(store CreateEmailStorage) *createEmailBiz {
	return &createEmailBiz{store: store}
}

func (biz *createEmailBiz) CreateNewEmail(ctx context.Context, data *modelemail.EmailCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateEmail(ctx, data); err != nil {
		return err
	}

	return nil
}
