package main

import(
	"fmt"
)

func main(){
	var arr [5]int
	for i:=0; i<len(arr); i++{
		arr[i]=i
	}
	fmt.Println(arr)
	
	arr2d := [3][2]int{{1,2},{2,3},{4,5}}
	for i:=0; i<len(arr2d); i++{
		for j:=0; j<len(arr2d[0]); j++{
			fmt.Println(arr2d[i][j])
		}
	}
	fmt.Println(arr2d)
}