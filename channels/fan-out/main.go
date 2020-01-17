/*
	Fanning out to run these in different subroutines
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1) // Puts values on channel, closes channel

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(v2 int) { // Fanning out all the values of c1 so launching 100 go routines
			c2 <- timeConsumingWork(v2) // Delayed response that returns new int
			wg.Done()
		}(v)
	}
	wg.Wait() // Wait for the group
	close(c2) // Close channel
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
