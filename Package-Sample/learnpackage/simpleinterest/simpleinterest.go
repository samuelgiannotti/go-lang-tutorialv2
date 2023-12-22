package simpleinterest

import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("Simple interest package initialized")
}

//Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
//Calculate is capitalized because it is accessed from outside the package. If it is renamend to calculate it will not work.
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
