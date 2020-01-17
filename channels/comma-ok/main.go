package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)
}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received in the even channel is:", v)
		case v := <-odd:
			fmt.Println("the value in the odd channel is:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			}

			fmt.Println("from comma ok", i)
		}
	}
}
