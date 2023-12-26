package main

import "fmt"

func main() {
	basicPointer()
	basicPointer2()
	pointerNilSample()
	arrayOfPointers()
	pointerToPointer()
	pointerToFunction()
}

func basicPointer() {
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a)
}

func basicPointer2() {
	fmt.Printf("basicPointer2 \n")
	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)
	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	a = 22
	fmt.Printf("New Value of *ip variable: %d\n", *ip)

}

func pointerNilSample() {
	fmt.Printf("pointerNilSample \n")
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)
}

const MAX int = 3

func arrayOfPointers() {
	fmt.Printf("arrayOfPointers \n")
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}

func pointerToPointer() {
	fmt.Printf("pointerToPointer \n")
	var a int
	var ptr *int
	var pptr **int

	a = 3000
	/* take the address of var */
	ptr = &a
	/* take the address of ptr using address of operator & */
	pptr = &ptr

	/* take the value using pptr */
	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Value available at **pptr = %d\n", **pptr)
}

func pointerToFunction() {
	fmt.Printf("pointerToFunction \n")
	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	* &a indicates pointer to a ie. address of variable a and
	* &b indicates pointer to b ie. address of variable b.
	 */
	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* save the value at address x */
	*x = *y   /* put y into x */
	*y = temp /* put temp into y */
}
