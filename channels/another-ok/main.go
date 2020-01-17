package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok1 := <-c
	w, ok2 := <-c // going to not have anything on here

	fmt.Println(v, ok1)
	fmt.Println(w, ok2)
}
