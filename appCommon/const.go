package appCommon

import "fmt"

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}

type TokenPayload struct {
	Uid   int    `json:"user_id"`
	Urole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.Uid
}

func (p TokenPayload) Role() string {
	return p.Urole
}
