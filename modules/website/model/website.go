package modelwebsite

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

const (
	EntityName = "website"
)

var (
	ErrWebsiteIsDeleted    = errors.New("website is deleted")
	ErrPathCannotBeEmpty   = errors.New("path cannot be empty")
	ErrLimitInvalid        = errors.New("limit is invalid")
	ErrRetryInvalid        = errors.New("retry is invalid")
	ErrEmailsCannotBeEmpty = errors.New("emails cannot be empty")
)

type Website struct {
	appCommon.SQLModel
	Path   string `json:"name" gorm:"column:name;"`
	Limit  int    `json:"limit" gorm:"column:limit;"`
	Retry  int    `json:"retry" gorm:"column:retry;"`
	Emails string `json:"emails" gorm:"column:emails;"`
	//Discords   string `json:"discords" gorm:"column:discords;"`
	//Facebooks  string `json:"facebooks" gorm:"column:facebooks;"`
	//Phones     string `json:"phones" gorm:"column:phones;"`
	//OtherLinks string `json:"other_links" gorm:"column:other_links;"`
	Status string `json:"status" gorm:"column:status;"`
}

func (Website) TableName() string {
	return "websites"
}

type WebsiteCreation struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Path   string `json:"name" gorm:"column:name;"`
	Limit  int    `json:"limit" gorm:"column:limit;"`
	Retry  int    `json:"retry" gorm:"column:retry;"`
	Emails string `json:"emails" gorm:"column:emails;"`
	//Discords   string `json:"discords,omitempty" gorm:"column:discords;"`
	//Facebooks  string `json:"facebooks,omitempty" gorm:"column:facebooks;"`
	//Phones     string `json:"phones,omitempty" gorm:"column:phones;"`
	//OtherLinks string `json:"other_links,omitempty" gorm:"column:other_links;"`
}

func (data *WebsiteCreation) Validate() error {
	data.Path = strings.TrimSpace(data.Path)
	if data.Path == "" {
		return ErrPathCannotBeEmpty
	}

	if data.Limit <= 0 || data.Limit > 24 {
		return ErrLimitInvalid
	}

	if data.Retry <= 0 || data.Retry > 24 {
		return ErrRetryInvalid
	}

	data.Emails = strings.TrimSpace(data.Emails)
	if data.Emails == "" {
		return ErrEmailsCannotBeEmpty
	}

	return nil
}

func (WebsiteCreation) TableName() string { return Website{}.TableName() }

type WebsiteUpdate struct {
	Path   *string `json:"name" gorm:"column:name;"`
	Limit  *int    `json:"limit" gorm:"column:limit;"`
	Retry  *int    `json:"retry" gorm:"column:retry;"`
	Emails *string `json:"emails" gorm:"column:emails;"`
	Status *string `json:"status" gorm:"column:status;"`
	//Discords   *string `json:"discords" gorm:"column:discords;"`
	//Facebooks  *string `json:"facebooks" gorm:"column:facebooks;"`
	//Phones     *string `json:"phones" gorm:"column:phones;"`
	//OtherLinks *string `json:"other_links" gorm:"column:other_links;"`
}

func (WebsiteUpdate) TableName() string { return Website{}.TableName() }

func (data *WebsiteUpdate) Validate() error {
	path := strings.TrimSpace(*data.Path)
	if path == "" {
		return ErrPathCannotBeEmpty
	}

	limit := *data.Limit
	if limit <= 0 || limit > 24 {
		return ErrLimitInvalid
	}

	retry := *data.Retry
	if retry <= 0 || retry > 24 {
		return ErrRetryInvalid
	}

	email := strings.TrimSpace(*data.Emails)
	if email == "" {
		return ErrEmailsCannotBeEmpty
	}

	return nil
}
