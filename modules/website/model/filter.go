package modelwebsite

type Filter struct {
	Status string `json:"status,omitempty" form:"status"`
}
