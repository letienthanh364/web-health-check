package modelwebsite

import (
	"errors"
	"strings"
)

const (
	WebsiteCheckTimeEntity = "website check time"
	CheckTimeLimit         = 5
)

var (
	ErrCheckTimeCannotBeEmpty = errors.New("check time cannot be empty")
	ErrCheckTimeIsExisted     = errors.New("check time is already existed")
	ErrCheckTimeExceedLimit   = errors.New("the number of check times is exceeding the limit")
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
