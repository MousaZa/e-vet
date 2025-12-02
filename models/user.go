package models

import (
	"fmt"
)

type User struct {
	ID            string         `json:"id"`
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	Notifications []Notification `json:"notifications"`
}

func (u *User) Update(productName string) {
	fmt.Printf("Running short on %s.", productName)
}

func (u *User) GetID() string {
	return u.ID
}
