package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	createdAt time.Time
}

func NewUser(firstname, lastname, birthdate string) (*User, error) {
	if firstname == "" || lastname == "" || birthdate == "" {
		return nil, errors.New("PAssed wrong values. Please check")
	}
	return &User{
		FirstName: firstname,
		LastName:  lastname,
		BirthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
func (u User) OutPutInfo() {
	fmt.Println(u.FirstName, " ", u.LastName, ", ", u.createdAt)
}

func (u *User) ClearName() {
	u.FirstName = ""
}
