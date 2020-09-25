package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	fmt.Println("Enter an integer: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := 0; i < int(number); i++{
		if i%3 == 0 && i%5 == 0{
			fmt.Println("fizzbuzz")
		}else if i%3 == 0{
			fmt.Println("fizz")
		}else if i%5 == 0{
			fmt.Println("buzz")
		}else{
			fmt.Println(i)
		}
	}
}