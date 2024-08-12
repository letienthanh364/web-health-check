package modelwebsite

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

const (
	WebsiteCheckTimeEntity = "website check time"
	CheckTimeLimit         = 5
)

var (
	ErrCheckTimeCannotBeEmpty = appCommon.NewErrorResponse(errors.New("check time cannot be empty"), "check time cannot be empty", "check time cannot be empty", "ErrCheckTimeCannotBeEmpty")
	ErrCheckTimeExisted       = appCommon.NewErrorResponse(errors.New("check time cannot be empty"), "check time is already existed", "check time is already existed", "ErrCheckTimeExisted")
	ErrCheckTimeExceedLimit   = appCommon.NewErrorResponse(errors.New("check time cannot be empty"), "the number of check times is exceeding the limit", "the number of check times is exceeding the limit", "ErrCheckTimeExceedLimit")
)

type WebsiteCheckTime struct {
	Id        int    `json:"id" gorm:"column:id;"`
	WebsiteId int    `json:"website_id" gorm:"column:website_id;"`
	CheckTime string `json:"check_time" gorm:"column:check_time;"`
}

func (WebsiteCheckTime) TableName() string {
	return "website_checktimes"
}

type WebsiteCheckTimeCreation struct {
	WebsiteId int    `json:"website_id" gorm:"column:website_id;"`
	CheckTime string `json:"check_time" gorm:"column:check_time;"`
}

func (WebsiteCheckTimeCreation) TableName() string { return WebsiteCheckTime{}.TableName() }

func (data *WebsiteCheckTimeCreation) Validate() error {
	checktime := strings.TrimSpace(data.CheckTime)
	if checktime == "" {
		return ErrCheckTimeCannotBeEmpty
	}

	return nil
}

type WebsiteCheckTimeDelete struct {
	Id int `json:"id" gorm:"column:id;"`
}

func (WebsiteCheckTimeDelete) TableName() string { return WebsiteCheckTime{}.TableName() }
