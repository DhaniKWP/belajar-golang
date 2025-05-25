package main

import	"fmt"

func main() {
	withDataType()
	withoutDataType()
	multipleVariable()
	underscoreVariable()
	printfUsage()
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
	// var name = "Dhani"
	// var age = 19

	//Short declaration
	name := "Dhani"
	age := 19

	fmt.Println("Hello", name, "Umur saya", age)
	fmt.Printf("%T, %T\n", name, age)
}

func multipleVariable() {
	//Multiple variable
	// var student1, student2, student3 string = "Joko", "lele", "Paus"
	// var angka1, angka2, angka3 int
	// angka1, angka2, angka3 = 1, 3, 5

	// fmt.Println(student1, student2, student3)
	// fmt.Println(angka1, angka2, angka3)

	var name, age, address = "Dhani", 19, "New York"
	first, second, thrid := "1", 2, "3"

	fmt.Println(name, age, address)
	fmt.Println(first, second, thrid)
}

func underscoreVariable() {
	//menanggulanginya adalah dengan memakai variable underscore
	var firstvariable string
	var name, age, address = "Blublu", 2, "cave"

	_,_,_,_ = firstvariable, name, age, address
}

func printfUsage() {

	// first, second := "1", 2

	// fmt.Printf("Tipe data dari variable pertama adalah %T\n", first)
	// fmt.Printf("Tipe data dari variable kedua adalah %T", second)

	//dapat menggunakan lebih dari 1 verb

	var nama = "Dhani"
	var umur = 19
	var alamat = "new york"

	fmt.Printf("nama saya adalah %s, umur saya %d, saya tinggal di %s", nama, umur, alamat)
}