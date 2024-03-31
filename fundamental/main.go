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

	// function
	singleParam("ibam")
	multiParam("ibrahim", "nurandita", 30)
	fmt.Println(multiParamWithReturn("ibrahim", "nurandita", 30))

	// conditions
	score := 10
	if score > 5 {
		fmt.Println("bagus")
	} else {
		fmt.Println("jelek")
	}

	// conditions dengan switch
	pakaian := "baju"
	switch pakaian {
	case "baju":
		fmt.Println("masuk ke rak baju")
	case "celana":
		fmt.Println("masuk ke rak celana")
	default:
		fmt.Println("buang")
	}

	// looping cara pertama
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// looping cara kedua
	// i := 0
	// for i < 3 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// looping range
	// i := []int{1, 2, 3} // slice of integer
	// for _, v := range i {
	// 	fmt.Println("Angka", v)
	// }

	// slice
	fruits := []string{"apel", "pisang", "jambu"}
	fmt.Println(fruits)
	// init slice dengan make
	fruits2 := make([]string, 3)
	fmt.Printf("%#v\n", fruits2)
	fruits2[0] = "pepaya"
	fruits2[1] = "mangga"
	fruits2[2] = "manggis"
	fmt.Printf("%#v\n", fruits2)
	// append slice
	fruits3 := append(fruits, fruits2...)
	fmt.Printf("%#v\n", fruits3)
}

// singleParam adalah contoh function single param
func singleParam(name string) {
	fmt.Println("halo", name)
}

// multiParam adalah contoh function multiParam param
func multiParam(firstName, lastName string, age int) {
	fmt.Println("halo", firstName, lastName, age)
}

// multiParamWithReturn adalah contoh function multiParam dengan return
func multiParamWithReturn(firstName, lastName string, age int) string {
	return fmt.Sprintf("firstName %s lastName %s age %d", firstName, lastName, age)
}
