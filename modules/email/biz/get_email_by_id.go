package bizemail

import (
	"context"
	modelemail "github.com/teddlethal/web-health-check/modules/email/model"
)

type GetEmailStorage interface {
	GetEmail(ctx context.Context, cond map[string]interface{}) (*modelemail.Email, error)
}

type getEmailBiz struct {
	store GetEmailStorage
}

func NewGetEmailBiz(store GetEmailStorage) *getEmailBiz {
	return &getEmailBiz{store: store}
}

func (biz *getEmailBiz) GetEmailById(ctx context.Context, id int) (*modelemail.Email, error) {
	data, err := biz.store.GetEmail(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	if data.Status == "deleted" {
		return nil, modelemail.ErrEmailIsDeleted
	}

	return data, nil
}
