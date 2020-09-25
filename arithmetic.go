package main

import (
	"fmt"
)

func main(){
	var num1 float32 = 10
	var num2 float32 = 3
	var num3 int = 10
	var num4 int = 3
	quotient := num1/num2 
	fmt.Println(quotient)
	quotientInt := num3/num4
	fmt.Printf("%v\n", quotientInt)
	fmt.Printf("%v", float32(quotientInt))
}