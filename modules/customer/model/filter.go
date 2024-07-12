package modelcustomer

type Filter struct {
	WebsiteId *int `json:"website_id,omitempty" gorm:"column:website_id;"`
}
