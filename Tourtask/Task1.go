package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	sum := 3.0
	for i := 0; (sum*sum - x)/(2*sum) > 0; i++ {
		sum = sum - (sum*sum - x)/(2*sum)
		fmt.Println(i)
	}

	return sum
}

func main() {
	fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}