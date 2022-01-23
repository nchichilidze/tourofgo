package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	
	var first, second, current int
	
	first = -1
	second = 1
	
	return func() int {
		current = first + second
		first = second
		second = current
		return current
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
