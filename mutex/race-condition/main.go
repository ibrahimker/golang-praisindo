package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	val int
}

func (c *Bank) Add(int) {
	c.val++
}

func (c *Bank) Value() int {
	return c.val
}

func main() {
	var wg sync.WaitGroup
	var bank Bank

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			bank.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("bank value", bank.Value())
}
