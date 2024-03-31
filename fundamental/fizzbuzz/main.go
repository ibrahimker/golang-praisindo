package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz", fmt.Sprintf("%d habis dibagi 3 dan 5", i))
		case i%3 == 0:
			fmt.Println("fizz", fmt.Sprintf("%d habis dibagi 3", i))
		case i%5 == 0:
			fmt.Println("buzz", fmt.Sprintf("%d habis dibagi 5", i))
		default:
			fmt.Println(i)
		}
	}
}
