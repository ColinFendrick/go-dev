package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("&x before", &x) // 0xc00008a008
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("&x after pointer fn", &x)
	fmt.Println("x after pointer fn", x)

	bar(x)
	fmt.Println("x after non-pointer fn does not get changed", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("*y before", *y) // 0xc00008c008
	*y = 43
	fmt.Println("y after pointer fn", y) // 0xc00008a008
	fmt.Println("*y after pointer fn", *y)
}

func bar(z int) {
	fmt.Println("z before non-pointer", z)
	fmt.Println("&z before non-pointer", &z) // 0xc00008a030 <- different memory address
	z = 100
	fmt.Println("z after non-pointer", z)
	fmt.Println("&z before non-pointer", &z)
}
