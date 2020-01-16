package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)
}

// send channel
func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// receive channel
func receive(c <-chan int) {
	for v := range c {
		fmt.Println("the value received from the channel:", v)
	}
}
