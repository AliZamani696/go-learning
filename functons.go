package main

import "fmt"


func add()int{
	var sum int =0
	for i :=0;i<50;i++{
		sum += i
		fmt.Println(i,sum)
	}
	return sum
}
func div(a,b int)(i,j int){
	i = a / b
	j = a % b
	return
}

func main(){

	fmt.Println(div(20,3))
	fmt.Println(20 / 3,20 % 3)
}