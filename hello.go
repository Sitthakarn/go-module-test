package hello

import (
	"fmt"

	"rsc.io/quote/v3"
)

// Hello function
func Hello() string {
	fmt.Println("Test hello program in ver 1.0.0")
	return quote.HelloV3()
}

// Proverb funtion
func Proverb() string {
	return quote.Concurrency()
}
