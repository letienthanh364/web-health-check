package modelwebsite

import (
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

const (
	EntityName             = "website"
	TimeIntervalLowerBound = 300
	RetryLowerBound        = 0
	RetryUpperBound        = 5
)

var (
	ErrWebsiteIsDeleted          = appCommon.NewErrorResponse(nil, "website is deleted", "website is deleted", "ErrWebsiteIsDeleted")
	ErrNameCannotBeEmpty         = appCommon.NewErrorResponse(nil, "name cannot be empty", "name cannot be empty", "ErrNameCannotBeEmpty")
	ErrPathCannotBeEmpty         = appCommon.NewErrorResponse(nil, "path cannot be empty", "path cannot be empty", "ErrPathCannotBeEmpty")
	ErrDefaultEmailCannotBeEmpty = appCommon.NewErrorResponse(nil, "default email cannot be empty", "default email cannot be empty", "ErrDefaultEmailCannotBeEmpty")
	ErrTimeIntervalInvalid       = appCommon.NewErrorResponse(nil, "time interval is too small", "time interval is too small", "ErrTimeIntervalInvalid")
	ErrRetryInvalid              = appCommon.NewErrorResponse(nil, "retry is invalid", "retry is invalid", "ErrRetryInvalid")
	ErrPathExisted               = appCommon.NewErrorResponse(nil, "website path is already existed", "website path is already existed", "ErrPathExisted")
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

	if data.TimeInterval < TimeIntervalLowerBound {
		return ErrTimeIntervalInvalid
	}

	if data.Retry < RetryLowerBound || data.Retry > RetryUpperBound {
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
		if timeInterval < TimeIntervalLowerBound {
			return ErrTimeIntervalInvalid
		}
	}

	if data.Retry != nil {
		retry := *data.Retry
		if retry < RetryLowerBound || retry > RetryUpperBound {
			return ErrRetryInvalid
		}
	}

	return nil
}
