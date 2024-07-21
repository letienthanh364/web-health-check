package modelemail

type Filter struct {
	Status    string `json:"status,omitempty" form:"status"`
	WebsiteId string `json:"website_id,omitempty" form:"website_id"`
}
