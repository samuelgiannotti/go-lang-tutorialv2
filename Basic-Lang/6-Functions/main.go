package main

import (
	"fmt"
	"math"
)

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200
	var ret int
	/* calling a function to get max value */
	ret = max(a, b)
	fmt.Printf("Max value is : %d\n", ret)
	c, d := swap("Mahesh", "Kumar")
	fmt.Println(c, d)

	//call by value
	/* local variable definition */
	fmt.Printf("call by value\n")
	var e int = 100
	var f int = 200
	fmt.Printf("Before swap, value of a : %d\n", e)
	fmt.Printf("Before swap, value of b : %d\n", f)
	/* calling a function to swap the values */
	callByValue(e, f)
	fmt.Printf("After swap, value of a : %d\n", e)
	fmt.Printf("After swap, value of b : %d\n", f)
	//call by reference
	/* local variable definition */
	fmt.Printf("call by reference\n")

	/* local variable definition */
	var g int = 100
	var h int = 200
	fmt.Printf("Before swap, value of a : %d\n", g)
	fmt.Printf("Before swap, value of b : %d\n", h)
	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	callByReference(&g, &h)
	fmt.Printf("After swap, value of a : %d\n", g)
	fmt.Printf("After swap, value of b : %d\n", h)

	//function as value
	fmt.Printf("function as value\n")
	functionAsValue()

	//anonymous functions or function closure
	/* nextNumber is now a function with i as 0 */
	fmt.Printf("function closure\n")
	nextNumber := getSequence()
	/* invoke nextNumber to increase i by 1 and return the same */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	/* create a new sequence and see the result, i is 0 again*/
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	//struct method definition
	fmt.Printf("method definition\n")
	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())
}

/* function returning the max between two numbers */
func max(num1, num2 int) int {
	/* local variable declaration */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//returing multiple values
func swap(x, y string) (string, string) {
	return y, x
}

func callByValue(x, y int) int {
	var temp int
	temp = x /* save the value of x */
	x = y    /* put y into x */
	y = temp /* put temp into y */
	return temp
}

func callByReference(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}

func functionAsValue() {
	/* declare a function variable */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* use the function */
	fmt.Println(getSquareRoot(9))
}

//function closure or anonymous functions
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//method
/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
