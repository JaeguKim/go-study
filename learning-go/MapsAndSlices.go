package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "fido"
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	myMap2["Second"] = 2
	log.Println(myMap2["First"],myMap2["Second"])

	myMap3 := make(map[string]User)
	me := User {
		FirstName: "Jaegoo",
		LastName: "Kim",
	}
	myMap3["me"] = me
	log.Println(myMap3["me"].FirstName)

	var mySlice []string
	mySlice = append(mySlice, "Trevor","John","Mary")
	log.Println(mySlice)
	var mySlice2 []int

	mySlice2 = append(mySlice2, 2,1,3)
	sort.Ints(mySlice2)
	log.Println(mySlice2)

	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	log.Println(numbers)
	log.Println(numbers[0:2])

	names := []string{"one","seven","fish","cat"}
	log.Println(names)
}
