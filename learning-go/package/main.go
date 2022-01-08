package main

import (
	"log"

	"github.com/jaegookim/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName,myVar.TypeNumber)
}