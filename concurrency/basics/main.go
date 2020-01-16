/*
But when people hear the word concurrency they often think of
parallelism, a related but quite distinct concept.
In programming, concurrency is the composition of
independently executing processes, while parallelism is the simultaneous execution of
(possibly related) computations.
Concurrency is about dealing with lots of things at once.
Parallelism is about doing lots of things at once.
*/

/*
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines
to wait for. Then each of the goroutines runs and calls
Done when finished. At the same time
Wait can be used to block until all goroutines have finished.
Writing concurrent code is super easy: all we do is put “go”
in front of a function or method call.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
