package main

import (
	"fmt"
)

func main() {
	/*
		// jika run sync
		result := introduceWithoutChannel("ibam")
		fmt.Println("hasil introduceWithoutChannel", result)

		// jika di run async, untuk dapat return value, gunakan channel
		c := make(chan string)
		defer close(c)

		go introduceWithChannel("ibam-async", c)
		msg1 := <-c
		fmt.Println("hasil introduceWithChannel", msg1)
	*/

	// buffered vs unbuffered
	// unbuffered channel
	// expected error karena channel blocking, perlu gunakan buffered channel
	/*
		c := make(chan string)
		c <- "Cat"
		message := <-c
		fmt.Println(message)
	*/
	// buffered channel
	/*
		c := make(chan string, 1)
		c <- "Cat"
		message := <-c
		fmt.Println(message)
	*/
}

func introduceWithoutChannel(name string) string {
	res := fmt.Sprintf("halo nama saya " + name)
	return res
}

func introduceWithChannel(name string, c chan string) {
	res := fmt.Sprintf("halo nama saya " + name)
	c <- res
}
