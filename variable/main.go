package main

import	"fmt"

func main() {
	withDataType()
	withoutDataType()
}

func withDataType() {
	// variable with data type
	var name string 
	var age int = 19

	name = "Dhani"
	age = 18

	fmt.Println("Nama saya Adalah", name)
	fmt.Println("Umur saya Adalah", age)
}

func withoutDataType() {
	// Variable without data type/type inference
	var name = "Dhani"
	var age = 19

	fmt.Printf("%T, %T", name, age)
}