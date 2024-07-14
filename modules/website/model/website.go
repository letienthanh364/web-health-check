package modelwebsite

const (
	EntityName = "website"
)

type Website struct {
	name string `json:"name" gorm:"column:name;"`
}

func (Website) TableName() string {
	return "websites"
}
