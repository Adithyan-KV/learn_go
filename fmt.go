package main

import "fmt"

func main() {
	var number int
	number = 300
	random := 100
	fmt.Printf("%T, %v\n", number, number)
	fmt.Printf("%T, %v\n", random, random)
	var x string = fmt.Sprintf("%T, %v", number, number)
	fmt.Println(x)
	var flt float64 = 324.2351253245234
	fmt.Printf("Scientific: %e, decimal: %f, large: %g", flt, flt, flt)
}
