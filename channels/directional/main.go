package main

import "fmt"

func main() {
	soc := make(chan<- int, 2)
	soc <- 42
	// fmt.Println(<-soc) will error out (receive from send-only)

	roc := make(<-chan int)
	fmt.Println(<-roc)

	// go func() {
	// 	roc <- 42 // send to receive-only channel
	// }()
}
