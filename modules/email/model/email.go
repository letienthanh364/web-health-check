package modelemail

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

const (
	EntityName = "website"
)

var (
	ErrEmailIsDeleted       = errors.New("email is deleted")
	ErrAddressCannotBeEmpty = errors.New("address cannot be empty")
)

type Email struct {
	appCommon.SQLModel
	Status  string `json:"status" gorm:"column:status;"`
	EmailId string `json:"website_id" gorm:"column:website_id;"`
	Address string `json:"address" gorm:"column:address;"`
}

func (Email) TableName() string {
	return "emails"
}

type EmailCreation struct {
	Id      int    `json:"id" gorm:"column:id;"`
	EmailId int    `json:"website_id" gorm:"column:website_id;"`
	Address string `json:"address" gorm:"column:address;"`
}

func (data *EmailCreation) Validate() error {
	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return ErrAddressCannotBeEmpty
	}

	return nil
}

func (EmailCreation) TableName() string { return Email{}.TableName() }

type EmailUpdate struct {
	Address *string `json:"address" gorm:"column:address;"`
	Status  *string `json:"status" gorm:"column:status;"`
}

func (EmailUpdate) TableName() string { return Email{}.TableName() }

func (data *EmailUpdate) Validate() error {
	address := strings.TrimSpace(*data.Address)
	if address == "" {
		return ErrAddressCannotBeEmpty
	}

	return nil
}
