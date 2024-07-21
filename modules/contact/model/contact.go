package modelcontact

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

const (
	EntityName = "contact"
)

var (
	ErrContactIsDeleted           = errors.New("contact is deleted")
	ErrWebsiteIdIsRequired        = errors.New("website id is required")
	ErrAddressCannotBeEmpty       = errors.New("address cannot be empty")
	ErrContactMethodCannotBeEmpty = errors.New("contact method cannot be empty")
)

type Contact struct {
	appCommon.SQLModel
	Status        string `json:"status" gorm:"column:status;"`
	WebsiteId     string `json:"website_id" gorm:"column:website_id;"`
	Address       string `json:"address" gorm:"column:address;"`
	ContactMethod string `json:"contact_method" gorm:"column:contact_method;"`
}

func (Contact) TableName() string {
	return "contacts"
}

type ContactCreation struct {
	Id            int    `json:"id" gorm:"column:id;"`
	WebsiteId     int    `json:"website_id" gorm:"column:website_id;"`
	Address       string `json:"address" gorm:"column:address;"`
	ContactMethod string `json:"contact_method" gorm:"column:contact_method;"`
}

func (data *ContactCreation) Validate() error {
	if data.WebsiteId == 0 {
		return ErrWebsiteIdIsRequired
	}

	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return ErrAddressCannotBeEmpty
	}

	data.ContactMethod = strings.TrimSpace(data.ContactMethod)
	if data.ContactMethod == "" {
		return ErrContactMethodCannotBeEmpty
	}

	return nil
}

func (ContactCreation) TableName() string { return Contact{}.TableName() }

type ContactUpdate struct {
	Status        *string `json:"status" gorm:"column:status;"`
	Address       *string `json:"address" gorm:"column:address;"`
	ContactMethod *string `json:"contact_method" gorm:"column:contact_method;"`
}

func (ContactUpdate) TableName() string { return Contact{}.TableName() }

func (data *ContactUpdate) Validate() error {

	address := strings.TrimSpace(*data.Address)
	if address == "" {
		return ErrAddressCannotBeEmpty
	}

	contactMethod := strings.TrimSpace(*data.ContactMethod)
	if contactMethod == "" {
		return ErrContactMethodCannotBeEmpty
	}

	return nil
}
