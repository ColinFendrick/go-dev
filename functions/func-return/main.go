package main

import "fmt"

func main() {
	x := bar("Hello ")
	fmt.Println(x("World"))

	fmt.Println(curry("curry ")("functions ")("are ")("delicious "))
}

func bar(s1 string) func(string) string {
	return func(s2 string) string {
		return s1 + s2
	}
}

func curry(s1 string) func(string) func(string) func(string) string {
	return func(s2 string) func(string) func(string) string {
		return func(s3 string) func(string) string {
			return func(s4 string) string {
				return s1 + s2 + s3 + s4
			}
		}
	}
}
