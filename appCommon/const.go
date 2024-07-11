package appCommon

import "fmt"

const (
	CurrentUser = "current_suer"
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}

type TokenPayload struct {
	Uid   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.Uid
}

func (p TokenPayload) Role() string {
	return p.URole
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin"
}
