package configmodel

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
)

var (
	milisecondsInADay  = 86400000
	lowerTime          = 1000 * 60 * 60
	upperTime          = 1000 * 60 * 60 * 24
	ErrConfigIsDeleted = errors.New("config is deleted")
	ErrInvalidTime     = errors.New("time before check must be from 1 hour to 24 hours")
	ErrInvalidLimit    = errors.New("the limit each checking must be from 1 to 5")
)

type Config struct {
	appCommon.SQLModel
	Time      int `json:"time" gorm:"column:time"`
	Limit     int `json:"limit" gorm:"column:limit"`
	StartTime int `json:"start_time" gorm:"column:start_time"`
}

func (Config) TableName() string { return "configs" }

type ConfigCreation struct {
	Id        int `json:"id" gorm:"column:id;"`
	Time      int `json:"time" gorm:"column:time"`
	Limit     int `json:"limit" gorm:"column:limit"`
	StartTime int `json:"start_time" gorm:"column:start_time"`
}

func (ConfigCreation) TableName() string { return Config{}.TableName() }

func (data *ConfigCreation) Validate() error {
	if data.Time < lowerTime || data.Time > upperTime {
		return ErrInvalidTime
	}

	if data.Limit < 1 || data.Limit > 5 {
		return ErrInvalidLimit
	}

	if data.StartTime < 0 || data.StartTime > milisecondsInADay {
		data.StartTime = 0
	}

	return nil
}

type ConfigUpdate struct {
	Time      *int    `json:"time" gorm:"column:time"`
	Limit     *int    `json:"limit" gorm:"column:limit"`
	StartTime *int    `json:"start_time" gorm:"column:start_time"`
	Status    *string `json:"status" gorm:"column:status;"`
}

func (data *ConfigUpdate) Validate() error {
	if *data.Time < lowerTime || *data.Time > upperTime {
		return ErrInvalidTime
	}

	if *data.Limit < 1 || *data.Limit > 5 {
		return ErrInvalidLimit
	}

	if *data.StartTime < 0 || *data.StartTime > milisecondsInADay {
		*data.StartTime = 0
	}

	return nil
}

func (ConfigUpdate) TableName() string { return Config{}.TableName() }
