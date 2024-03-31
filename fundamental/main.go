package main

import "fmt"

func main() {
	// variable dengan format var <nama variabel> <tipe data> = <nilai>
	var firstName string = "john"
	// bisa juga ditulis seperti ini
	var lastName string
	lastName = "wick"
	// contoh inline variable
	middleName := "ahmad"
	// contoh multi variable
	first, second, third := "satu", "dua", "tiga"
	// contoh underscore variable
	_ = "tidak terpakai"

	fmt.Println(firstName, middleName, lastName)   // dengan println
	fmt.Printf("%s %s %s\n", first, second, third) // dengan printf

	contohBool := true
	contohInt := 1
	contohFloat := 1.5
	var contohDefaultInt int

	fmt.Printf("bool: [%t] int: [%d] float 2 angka belakang koma: [%.2f] default int: [%d]\n",
		contohBool, contohInt, contohFloat, contohDefaultInt)
}
