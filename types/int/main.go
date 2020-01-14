package main

import (
	"fmt"
	"runtime"
)

var x int8 // -128 to 127
var y float32

func main() {
	x = 127
	y = 42.34534234987249832749832794
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	// Running program os (darwin)
	fmt.Println(runtime.GOOS)
	// Architecture (amd64)
	fmt.Println(runtime.GOARCH)
}
