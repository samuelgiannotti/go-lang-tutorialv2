package main

import "fmt"

func main() {
	fmt.Printf("Hello\tWorld!\n")

	//int
	const x1 int = 212
	const x2 uint64 = 215
	const x3 int = 0xFee
	fmt.Printf("value of x1 : %d\n", x1)
	fmt.Printf("value of x2 : %d\n", x2)
	fmt.Printf("value of x3 : %d\n", x3)

	//float point
	var x4 float64 = 3.14159
	fmt.Printf("value of x4 : %f\n", x4)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)

	var f float32 = 1
	var i int = 1.000
	var u uint32 = 1e3 - 99.0*10.0 - 9
	var c float64 = '\x01'
	var p uintptr = '\u0001'
	var r complex64 = 'b' - 'a'
	var b byte = 1.0 + 3i - 3.0i

	fmt.Printf("value of f : %f\n", f)
	fmt.Printf("value of i : %d\n", i)
	fmt.Printf("value of u : %d\n", u)
	fmt.Printf("value of c : %f\n", c)
	fmt.Printf("value of p : %d\n", p)
	fmt.Println("value of r : \n", r)
	fmt.Printf("value of b : %d\n", b)
}
