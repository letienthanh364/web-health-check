package modelconfig

type Filter struct {
	Time       *int `json:"time,omitempty" gorm:"column:time;"`
	CheckLimit *int `json:"check_limit,omitempty" gorm:"column:check_limit;"`
	StartTime  *int `json:"start_time,omitempty" gorm:"column:start_time;"`
}
