package main

import (
	"fmt"
	"sync"
)

func main() {
	// format run async di golang
	// go <function>

	// tidak akan muncul karena program selesai lebih dahulu sebelum di print
	// anonymous function - run asynchronously
	// go func() {
	// 	fmt.Println("hello world secara async")
	// }()

	// named function - run asynchronously
	// go runAsync()

	// menambahkan sleep 10 ms akan membuat resultnya muncul
	// time.Sleep(10 * time.Millisecond)

	// async process
	/*
		fmt.Println("main execution started")

		go firstProcess(8)

		secondProcess(8)

		fmt.Println("number of goroutines", runtime.NumGoroutine())
		fmt.Println("main execution ended")

		// menambahkan sleep 10 ms akan membuat resultnya muncul
		time.Sleep(10 * time.Millisecond)
	*/

	// wait group
	/*
		fmt.Println("main execution started")

		var wg sync.WaitGroup
		wg.Add(1)
		go firstProcessWithWaitgroup(8, &wg)
		wg.Wait()

		secondProcess(8)

		fmt.Println("number of goroutines", runtime.NumGoroutine())
		fmt.Println("main execution ended")
	*/

	// looping async
	/*
		var wg sync.WaitGroup
		wg.Add(9)
		for i := 1; i < 10; i++ {
			go func() {
				fmt.Println(i)
				wg.Done()
			}()
		}
		wg.Wait()
	*/

}

func runAsync() {
	fmt.Println("hello world secara async")
}

func firstProcess(index int) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
}

func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second process func ended")
}

func firstProcessWithWaitgroup(index int, wg *sync.WaitGroup) {
	fmt.Println("first process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("first process func ended")
	wg.Done()
}
