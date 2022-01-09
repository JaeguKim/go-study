package main

import (
	"log"
)

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}
	animals := []string{"dog", "fish", "horse", "cow","cat"}
	for _, animal := range animals {
		log.Println(animal)
	}
	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"

	for animalType,animal := range animals2 {
		log.Println(animalType,animal)
	}

	var firstLine = "Once upon a midnight dreary"
	firstLine = "x"
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	var users []User
	users = append(users, User{"John","Smith","john@smith.com",30})
	users = append(users, User{"Mary","Jones","mary@smith.com",30})
	users = append(users, User{"Sally","Brown","sally@smith.com",30})
	users = append(users, User{"Alex","Anderson","alex@smith.com",30})
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
