package main

import (
	"fmt"
)

func main() {

	func(start int, end int, skip int) {
		for i := start; i < end; i += skip {
			fmt.Println(i)
		}
	}(4, 19, 2)
}
