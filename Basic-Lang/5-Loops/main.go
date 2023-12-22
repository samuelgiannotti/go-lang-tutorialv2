package main

import "fmt"

func main() {
	forLoopSample()
	nestedLoopSample()
	loopContinueSample()
	forGotoSample()
}

func forLoopSample() {
	var b int = 15
	var a int
	numbers := [6]int{1, 2, 3, 5}

	/* for loop execution */
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	fmt.Printf("after 1st loop value of a: %d\n", a)
	for a < b {
		a++
		fmt.Printf("for a<b value of a: %d\n", a)
	}
	fmt.Printf("after 2nd loop value of a: %d\n", a)
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}
}

func nestedLoopSample() {
	/* local variable definition */
	fmt.Printf("nestedLoop finding prime numbers\n")
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				fmt.Printf("%d factor found, not prime \n", i)
				break // if factor found, not prime
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}

func loopContinueSample() {
	fmt.Printf("loopContinueSample \n")
	/* local variable definition */
	var a int = 10

	/* do loop execution */
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}

func forGotoSample() {
	fmt.Printf("forGotoSample \n")
	/* local variable definition */
	var a int = 10

	/* do loop execution */
LOOP:
	for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}

func infiniteLoopSample() {
	for true {
		fmt.Printf("This loop will run forever.\n")
	}
}
