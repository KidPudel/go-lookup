package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println(sumTwo(1, 2))
	printVariables()
	converter()
}

func sumTwo(a, b int) (int, int, int) {
	return a, b, a + b
}

func namedReturn(a int) (x, b int) {
	x = a + 1
	b = x + 1
	return
}

var c, python, golang = "c makes you a great developer", "python productive", "go efficient"

func printVariables() {
	i := "short declaration"
	fmt.Println(i, c, python, golang)
}

var a, b int = 1, 2

// a
func converter() {
	value := 10
	newThing := fmt.Sprint(value)
	newThing += " sisters"
	fmt.Println(newThing)

	number := 5.71
	newNumber := int(number)
	fmt.Println(newNumber)

}

func numericConstants() {
	const (
		Big   = 1 << 100
		Small = Big >> 96
	)
	fmt.Println(float64(Big*0.1), Small)
}

const Pi = 3.1415

func example(a, b int) (x, y int) {
	x = 1
	y = 2
	return x, y
}
