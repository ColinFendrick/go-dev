package mystr

import "strings"

// Cat is cat
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Join joins two strings around a space
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
