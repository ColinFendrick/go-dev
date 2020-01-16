package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Go routine one started")
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("Go routine two started")
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
		wg.Done()
	}()

	wg.Wait()
}
