package main

import (
	"log"
	"time"
)

// Go does not have access modifier
// captial variable can be accessed outside of package
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

var Special string

func main() {
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "1 555-555-1212",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}