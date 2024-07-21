package usermodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/teddlethal/web-health-check/appCommon"
	"net/http"
)

const EntityName = "User"

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value ", value))
	}

	var r UserRole

	roleValue := string(bytes)

	if roleValue == "user" {
		r = RoleUser
	} else if roleValue == "admin" {
		r = RoleAdmin
	}

	*role = r

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}

func (role *UserRole) MarshalJson() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

type User struct {
	appCommon.SQLModel
	Email     string   `json:"contact" gorm:"column:contact;"`
	Password  string   `json:"password" gorm:"column:password;"`
	Salt      string   `json:"salt" gorm:"column:salt;"`
	LastName  string   `json:"last_name" gorm:"column:last_name;"`
	FirstName string   `json:"first_name" gorm:"column:first_name;"`
	Phone     string   `json:"phone" gorm:"column:phone;"`
	Role      UserRole `json:"role" gorm:"column:role;"`
	Status    int      `json:"status" gorm:"column:status"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (u User) TableName() string {
	return "users"
}

type UserCreate struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Email     string `json:"contact" gorm:"column:contact;"`
	Password  string `json:"password" gorm:"column:password;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"-" gorm:"column:role;"`
	Salt      string `json:"-" gorm:"column:salt;"`
}

func (u UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"contact" gorm:"column:contact;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserLogin) TabletName() string {
	return User{}.TableName()
}

var (
	ErrEmailOrPasswordInvalid = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("contact or password invalid"),
		"contact or password invalid",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("contact has already existed"),
		"contact has already existed",
		"ErrEmailExist",
	)
)
