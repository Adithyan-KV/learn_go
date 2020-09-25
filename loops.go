package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	for i := 0; i < int(number); i++{
		fmt.Println(i)
	}
}