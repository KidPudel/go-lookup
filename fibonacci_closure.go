package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current := 0
	previous := 0
	return func () int {
		sum := current + previous
		if previous == 0 {
			previous = 1
		} else {
			previous = current
			current = sum
		}
		
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
