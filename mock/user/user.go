package user

import (
	"code.byted.org/hc_test/mock/person"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}
func (u *User) GetUserInfo(id int) int {
	if id < 0 {
		return -1
	} else {
		return u.Person.Get(id)
	}
}

func (u *User) GetUserName(idx int) string {
	return person.GetName(idx)
}
