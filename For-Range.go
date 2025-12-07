package main

import "fmt"

func main(){
	numbers := []int{10,20,30}
	for _,v := range numbers{
		if v < 20{
			fmt.Println(v)
		}
	}
	for i,v := range numbers{
		fmt.Printf("index: %d value %d\n",i,v)
	}
}


var i8 int8 = 126

var i16 int16 = 12334
var i32 int32 = 90