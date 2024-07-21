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
	ErrWebsiteIsDeleted           = errors.New("website is deleted")
	ErrNameCannotBeEmpty          = errors.New("name cannot be empty")
	ErrPathCannotBeEmpty          = errors.New("path cannot be empty")
	ErrLimitInvalid               = errors.New("limit is invalid")
	ErrRetryInvalid               = errors.New("retry is invalid")
	ErrDefaultEmailCannotBeEmpty  = errors.New("default_email cannot be empty")
	ErrContactLinkCannotBeEmpty   = errors.New("contact link cannot be empty")
	ErrContactMethodCannotBeEmpty = errors.New("contact method cannot be empty")
)

type Website struct {
	appCommon.SQLModel
	Name         string `json:"name" gorm:"column:name;"`
	Path         string `json:"path" gorm:"column:path;"`
	Limit        int    `json:"limit" gorm:"column:limit;"`
	Retry        int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail string `json:"default_email" gorm:"column:default_email;"`
	Status       string `json:"status" gorm:"column:status;"`
	//Discords   []int `json:"discords"`
	//Facebooks  []int `json:"facebooks"`
	//Phones     []int `json:"phones"`
	//OtherLinks []int `json:"other_links"`
}

func (Website) TableName() string {
	return "websites"
}

type WebsiteCreation struct {
	Id           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	Path         string `json:"path" gorm:"column:path;"`
	Limit        int    `json:"limit" gorm:"column:limit;"`
	Retry        int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail string `json:"default_email" gorm:"column:default_email;"`
	//Discords   string `json:"discords,omitempty" gorm:"column:discords;"`
	//Facebooks  string `json:"facebooks,omitempty" gorm:"column:facebooks;"`
	//Phones     string `json:"phones,omitempty" gorm:"column:phones;"`
	//OtherLinks string `json:"other_links,omitempty" gorm:"column:other_links;"`
}

func (data *WebsiteCreation) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameCannotBeEmpty
	}

	data.Path = strings.TrimSpace(data.Path)
	if data.Path == "" {
		return ErrPathCannotBeEmpty
	}

	data.DefaultEmail = strings.TrimSpace(data.DefaultEmail)
	if data.DefaultEmail == "" {
		return ErrDefaultEmailCannotBeEmpty
	}

	if data.Limit <= 0 || data.Limit > 43200 {
		return ErrLimitInvalid
	}

	if data.Retry <= 0 || data.Retry > 10 {
		return ErrRetryInvalid
	}

	return nil
}

func (WebsiteCreation) TableName() string { return Website{}.TableName() }

type WebsiteUpdate struct {
	Name         *string `json:"name" gorm:"column:name;"`
	Path         *string `json:"path" gorm:"column:path;"`
	Limit        *int    `json:"limit" gorm:"column:limit;"`
	Retry        *int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail *string `json:"default_email" gorm:"column:default_email;"`
	Status       *string `json:"status" gorm:"column:status;"`
	//Discords   *string `json:"discords" gorm:"column:discords;"`
	//Facebooks  *string `json:"facebooks" gorm:"column:facebooks;"`
	//Phones     *string `json:"phones" gorm:"column:phones;"`
	//OtherLinks *string `json:"other_links" gorm:"column:other_links;"`
}

func (WebsiteUpdate) TableName() string { return Website{}.TableName() }

func (data *WebsiteUpdate) Validate() error {
	name := strings.TrimSpace(*data.Path)
	if name == "" {
		return ErrNameCannotBeEmpty
	}

	path := strings.TrimSpace(*data.Path)
	if path == "" {
		return ErrPathCannotBeEmpty
	}

	email := strings.TrimSpace(*data.DefaultEmail)
	if email == "" {
		return ErrDefaultEmailCannotBeEmpty
	}

	limit := *data.Limit
	if limit <= 0 || limit > 24 {
		return ErrLimitInvalid
	}

	retry := *data.Retry
	if retry <= 0 || retry > 24 {
		return ErrRetryInvalid
	}

	return nil
}
