package main

import "fmt"

func main() {
	var x, y int
	m := min(x) // m == x
	fmt.Println(m)
	m = min(x, y) // m is the smaller of x and y
	fmt.Println(m)
	m = max(x, y, 10) // m is the larger of x and y but at least 10
	fmt.Println(m)
	c := max(1, 2.0, 10) // c == 10.0 (floating-point kind)
	fmt.Println(c)
	f := max(0, float32(x)) // type of f is float32
	fmt.Println(f)
	t := max("", "foo", "bar") // t == "foo" (string kind)
	fmt.Println(t)
}
