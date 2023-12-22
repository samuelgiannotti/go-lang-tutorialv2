package main

import "fmt"

/* global variable declaration */
var g int
var g1 int = 20

func main() {
	/* local variable declaration */
	var a, b, c int

	/* actual initialization */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)

	g = a + b
	fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)

	/* local variable declaration */
	var g int = 10

	fmt.Printf("value of g = %d\n", g)

	/* local variable declaration in main function */
	var g1 int = 10
	var d int = 20
	var e int = 0
	fmt.Printf("value of g1 in main() = %d\n", g1)
	e = sum(g, d)
	fmt.Printf("value of e in main() = %d\n", e)
}

/* function to add two integers */
func sum(g1, d int) int {
	fmt.Printf("value of g1 in sum() = %d\n", g1)
	fmt.Printf("value of d in sum() = %d\n", d)
	return g1 + d
}
