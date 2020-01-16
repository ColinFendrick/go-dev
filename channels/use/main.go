/*
create general channels
in funcs you can specify
receive channel
you can receive values from the channel
a receive channel parameter
in the func, you can only pull values from the channel
you can’t close a receive channel
send channel
you can push values to the channel
you can’t receive/pull/read from the channel
you can only push values to the channel
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)

	receive(c)
	close(c) // WE CLOSED THE CHANNEL
}

// send channel
func send(c chan<- int) {
	c <- 42
}

// receive channel
func receive(c <-chan int) {
	fmt.Println("the value received from the channel:", <-c)
}
