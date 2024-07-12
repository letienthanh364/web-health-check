package modelcustomer

import "github.com/teddlethal/web-health-check/appCommon"

type Customer struct {
	appCommon.SQLModel
	WebsiteId         int    `json:"website_id" gorm:"column:website_id;"`
	Name              string `json:"name" gorm:"column:name;"`
	Email             string `json:"email" gorm:"column:email;"`
	Phone             string `json:"phone" gorm:"column:phone;"`
	CommunicateMethod string `json:"communicate_method" gorm:"column:communicate_method;"`
	Link              string `json:"link" gorm:"column:link;"`
}

func (Customer) TableName() string {
	return "customers"
}

type CustomerCreate struct {
	Id                int    `json:"id" gorm:"column:id;"`
	WebsiteId         int    `json:"website_id,omitempty" gorm:"column:website_id;"`
	Name              string `json:"name" gorm:"column:name;"`
	Email             string `json:"email" gorm:"column:email;"`
	Phone             string `json:"phone" gorm:"column:phone;"`
	CommunicateMethod string `json:"communicate_method" gorm:"column:communicate_method;"`
	Link              string `json:"link" gorm:"column:link;"`
}

func (CustomerCreate) TableName() string {
	return "customers"
}

type CustomerUpdate struct {
	Id                *int    `json:"id" gorm:"column:id;"`
	WebsiteId         *int    `json:"website_id" gorm:"column:website_id;"`
	Name              *string `json:"name" gorm:"column:name;"`
	Email             *string `json:"email" gorm:"column:email;"`
	Phone             *string `json:"phone" gorm:"column:phone;"`
	CommunicateMethod *string `json:"communicate_method" gorm:"column:communicate_method;"`
	Link              *string `json:"link" gorm:"column:link;"`
}

func (CustomerUpdate) TableName() string {
	return "customers"
}
