package main

import(
	"fmt"
)

func main(){
	var a [5]int = [5]int{10,20,30,40,50}
	for i, element := range(a){
		fmt.Println(i)
		fmt.Println(element)
	}
	for _, element := range(a){
		fmt.Println(element)
	}
}