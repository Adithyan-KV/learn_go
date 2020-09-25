package main

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main(){
	fmt.Println("How old are you?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	if age >= 21{
		fmt.Println("you can drink")
	}else{
		fmt.Println("You're not old enough to drink")
		fmt.Printf("Wait %v years", 21-age)
	}
}
