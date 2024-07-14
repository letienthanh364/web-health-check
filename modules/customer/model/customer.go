package modelcustomer

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

var (
	ErrWebsiteIdCannotBeEmpty     = errors.New("website id cannot be empty")
	ErrNameCannotBeEmpty          = errors.New("name cannot be empty")
	ErrEmailCannotBeEmpty         = errors.New("email cannot be empty")
	ErrPhoneCannotBeEmpty         = errors.New("phone cannot be empty")
	ErrContactMethodCannotBeEmpty = errors.New("contact method cannot be empty")
	ErrLinkCannotBeEmpty          = errors.New("link cannot be empty")
	ErrWebsiteIdIsInvalid         = errors.New("website id is invalid")
	ErrCustomerIsDeleted          = errors.New("customer is deleted")
)

const (
	EntityName = "customer"
)

type Customer struct {
	appCommon.SQLModel
	WebsiteId     int    `json:"website_id" gorm:"column:website_id;"`
	Name          string `json:"name" gorm:"column:name;"`
	Email         string `json:"email" gorm:"column:email;"`
	Phone         string `json:"phone" gorm:"column:phone;"`
	ContactMethod string `json:"contact_method" gorm:"column:contact_method;"`
	Link          string `json:"link" gorm:"column:link;"`
	Status        string `json:"status" gorm:"column:status;"`
}

func (Customer) TableName() string {
	return "customers"
}

type CustomerCreate struct {
	Id            int    `json:"id" gorm:"column:id;"`
	WebsiteId     int    `json:"website_id,omitempty" gorm:"column:website_id;"`
	Name          string `json:"name" gorm:"column:name;"`
	Email         string `json:"email" gorm:"column:email;"`
	Phone         string `json:"phone" gorm:"column:phone;"`
	ContactMethod string `json:"contact_method" gorm:"column:contact_method;"`
	Link          string `json:"link" gorm:"column:link;"`
}

func (c *CustomerCreate) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return ErrNameCannotBeEmpty
	}

	c.Email = strings.TrimSpace(c.Email)
	if c.Email == "" {
		return ErrEmailCannotBeEmpty
	}

	c.Phone = strings.TrimSpace(c.Phone)
	if c.Phone == "" {
		return ErrPhoneCannotBeEmpty
	}

	c.ContactMethod = strings.TrimSpace(c.ContactMethod)
	if c.ContactMethod == "" {
		return ErrContactMethodCannotBeEmpty
	}

	c.Link = strings.TrimSpace(c.Link)
	if c.Link == "" {
		return ErrLinkCannotBeEmpty
	}

	return nil
}

func (CustomerCreate) TableName() string {
	return "customers"
}

type CustomerUpdate struct {
	Id            *int    `json:"id" gorm:"column:id;"`
	Status        *string `json:"status" gorm:"column:status;"`
	WebsiteId     *int    `json:"website_id" gorm:"column:website_id;"`
	Name          *string `json:"name" gorm:"column:name;"`
	Email         *string `json:"email" gorm:"column:email;"`
	Phone         *string `json:"phone" gorm:"column:phone;"`
	ContactMethod *string `json:"contact_method" gorm:"column:contact_method;"`
	Link          *string `json:"link" gorm:"column:link;"`
}

func (CustomerUpdate) TableName() string {
	return "customers"
}
