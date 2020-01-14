package main

import "fmt"

// Page ...
type Page interface {
	PrintPage()
}

// HTMLPage ...
type HTMLPage struct {
	// Implement Page interface.
	Page
}

// Image ...
type Image interface {
	PrintImage()
}

// ImagePage ...
type ImagePage struct {
	// Implement Image interface.
	Image
}

func test(value interface{}) {
	// Use type switch to test interface type.
	// ... The argument is an interface.
	switch value.(type) {
	case nil:
		fmt.Println("Is nil interface")
	case Page:
		fmt.Println("Is page interface")
	case Image:
		fmt.Println("Is image interface")
	}
}

func main() {
	// Create class that implements interface and pass to test func.
	item1 := new(HTMLPage)
	test(item1)

	item2 := new(ImagePage)
	test(item2)

	test(nil)
}
