package main

import (
	"fmt"
)

type ErrNegativeSqrt float64 

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}


func Sqrt(approx float64) (float64, error) {
	if approx < 0 { 
		return -1, ErrNegativeSqrt(approx)
	}
	var result float64
	var i int
	
	result = float64(1)
	for i = 1; i <= 10; i++ { 
		result -= (result * result - approx) / (2 * result)
	}
	
	return result, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
