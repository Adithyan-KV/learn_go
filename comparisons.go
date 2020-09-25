package main

import "fmt"

func main(){
	var num1 int = 10
	var num2 int = 10
	fmt.Printf("%t\n", num1<num2)
	fmt.Printf("%t\n", num1>num2)
	fmt.Printf("%t\n", num1<=num2)
	fmt.Printf("%t\n", num1>=num2)
	fmt.Printf("%t\n", num1==num2)
	fmt.Printf("%t\n", num1!=num2)
}