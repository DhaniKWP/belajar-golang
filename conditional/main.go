package main

import "fmt"

//IF ELSE

func IfELse() {
	var tahunsekarang = 2025

	if umur := tahunsekarang - 2006; umur < 17 {
		fmt.Println("Kamu belum boleh membuat kartu SIM")
		} else {
		fmt.Println("Kamu boleh membuat kartu SIM")
	}
}


func main() {
	IfELse()
}