package tokenprovider

import (
	"errors"
	"github.com/teddylethal/todo-list/appCommon"
	"net/http"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() int
	Role() string
}
type Token interface {
	GetToken() string
}

var (
	ErrNotFound = appCommon.NewCustomError(
		http.StatusNotFound,
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = appCommon.NewCustomError(
		http.StatusInternalServerError,
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)
