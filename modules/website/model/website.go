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
	TimeInterval int    `json:"time_interval" gorm:"column:time_interval;"`
	Retry        int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail string `json:"default_email" gorm:"column:default_email;"`
	Status       string `json:"status" gorm:"column:status;"`
}

func (Website) TableName() string {
	return "websites"
}

type WebsiteDetail struct {
	appCommon.SQLModel
	Name         string `json:"name" gorm:"column:name;"`
	Path         string `json:"path" gorm:"column:path;"`
	TimeInterval int    `json:"time_interval" gorm:"column:time_interval;"`
	Retry        int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail string `json:"default_email" gorm:"column:default_email;"`
	Status       string `json:"status" gorm:"column:status;"`
}

type WebsiteCreation struct {
	Id           int    `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	Path         string `json:"path" gorm:"column:path;"`
	TimeInterval int    `json:"time_interval" gorm:"column:time_interval;"`
	Retry        int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail string `json:"default_email" gorm:"column:default_email;"`
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

	if data.TimeInterval <= 0 || data.TimeInterval > 43200 {
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
	TimeInterval *int    `json:"time_interval" gorm:"column:time_interval;"`
	Retry        *int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail *string `json:"default_email" gorm:"column:default_email;"`
	Status       *string `json:"status" gorm:"column:status;"`
}

func (WebsiteUpdate) TableName() string { return Website{}.TableName() }

func (data *WebsiteUpdate) Validate() error {
	if data.Name != nil {
		name := strings.TrimSpace(*data.Name)
		if name == "" {
			return ErrNameCannotBeEmpty
		}
	}

	if data.Path != nil {
		path := strings.TrimSpace(*data.Path)
		if path == "" {
			return ErrPathCannotBeEmpty
		}
	}

	if data.DefaultEmail != nil {
		email := strings.TrimSpace(*data.DefaultEmail)
		if email == "" {
			return ErrDefaultEmailCannotBeEmpty
		}
	}

	if data.TimeInterval != nil {
		limit := *data.TimeInterval
		if limit <= 0 || limit > 24 {
			return ErrLimitInvalid
		}
	}

	if data.Retry != nil {
		retry := *data.Retry
		if retry <= 0 || retry > 24 {
			return ErrRetryInvalid
		}
	}

	return nil
}
