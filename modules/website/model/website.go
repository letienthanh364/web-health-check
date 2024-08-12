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
	ErrWebsiteIsDeleted          = errors.New("website is deleted")
	ErrNameCannotBeEmpty         = errors.New("name cannot be empty")
	ErrPathCannotBeEmpty         = errors.New("path cannot be empty")
	ErrTimeIntervalInvalid       = errors.New("time interval is invalid")
	ErrRetryInvalid              = errors.New("retry is invalid")
	ErrDefaultEmailCannotBeEmpty = errors.New("default_email cannot be empty")
	ErrPathIsExisted             = errors.New("website path is already existed")
)

type Website struct {
	appCommon.SQLModel
	Status string `json:"status" gorm:"column:status;"`

	Name         string `json:"name" gorm:"column:name;"`
	Path         string `json:"path" gorm:"column:path;"`
	TimeInterval int    `json:"time_interval" gorm:"column:time_interval;"`
	Retry        int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail string `json:"default_email" gorm:"column:default_email;"`
	//TimeZone     string `json:"time_zone" gorm:"column:time_zone;"`
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
	//TimeZone     string `json:"time_zone" gorm:"column:time_zone;"`
}

func (data *WebsiteCreation) Validate() error {
	name := strings.TrimSpace(data.Name)
	if name == "" {

		return ErrNameCannotBeEmpty
	}

	path := strings.TrimSpace(data.Path)
	if path == "" {
		return ErrPathCannotBeEmpty
	}

	defaultEmail := strings.TrimSpace(data.DefaultEmail)
	if defaultEmail == "" {
		return ErrDefaultEmailCannotBeEmpty
	}

	if data.TimeInterval <= 60 {
		return ErrTimeIntervalInvalid
	}

	if data.Retry < 0 || data.Retry > 10 {
		return ErrRetryInvalid
	}

	return nil
}

func (WebsiteCreation) TableName() string { return Website{}.TableName() }

type WebsiteUpdate struct {
	Status       *string `json:"status" gorm:"column:status;"`
	Name         *string `json:"name" gorm:"column:name;"`
	Path         *string `json:"path" gorm:"column:path;"`
	TimeInterval *int    `json:"time_interval" gorm:"column:time_interval;"`
	Retry        *int    `json:"retry" gorm:"column:retry;"`
	DefaultEmail *string `json:"default_email" gorm:"column:default_email;"`
	//TimeZone     *string `json:"time_zone" gorm:"column:time_zone;"`
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
		timeInterval := *data.TimeInterval
		if timeInterval <= 60 {
			return ErrTimeIntervalInvalid
		}
	}

	if data.Retry != nil {
		retry := *data.Retry
		if retry < 0 || retry > 24 {
			return ErrRetryInvalid
		}
	}

	return nil
}
