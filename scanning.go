package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	fmt.Println("Enter a number: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("type: %T\n", input)
	fmt.Printf("The number-10 is %v", input-10)
}