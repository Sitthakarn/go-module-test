package hello

import (
	"fmt"
	"math"

	"rsc.io/quote/v3"
)

// Hello function
func Hello() string {
	fmt.Println("Test hello program in ver 1.0.0")
	fmt.Printf("This is Pi = %v", math.Pi)
	return quote.HelloV3()
}

// Proverb funtion
func Proverb() string {
	return quote.Concurrency()
}
