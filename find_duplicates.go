package main

import(
	"fmt"
)

func main(){
	var arr [8]int = [8]int{1,2,3,4,3,7,8,4}

	for i, element:= range arr{
		for j:=i+1; j<len(arr);j++{
			element2 := arr[j]
			if element == element2{
				fmt.Println(element)
			}
		} 
	}
}