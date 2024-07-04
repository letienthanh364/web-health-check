package appCommon

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Status    string     `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
