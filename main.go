package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "S-Tech",
		LastName:    "Lab",
		PhoneNumber: "0123",
	}
	fmt.Println(user.FirstName, user.LastName, user.PhoneNumber, user.BirthDate)
}
