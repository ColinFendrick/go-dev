/*
	Error handling, rather than using try-catch-finally control structures, errors can get returned
	right in the function return. Error is a built-in type.
	The recovery mechanism is executed only as part of a function's
	state being torn down after an error, which is sufficient to handle catastrophe but requires no
	extra control structures and, when used well, can result in clean error-handling code.
*/

package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// Do one of the following
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}
