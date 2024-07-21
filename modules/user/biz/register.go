package biz

import (
	"context"
	"github.com/teddlethal/web-health-check/appCommon"
	usermodel "github.com/teddlethal/web-health-check/modules/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(storage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: storage,
		hasher:          hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"contact": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := appCommon.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return appCommon.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
