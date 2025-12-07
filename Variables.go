package main
import "fmt"

var Global = "this is global var"

func main(){
	var print = fmt.Println
	var booli bool = true
	var num int = 10
	num2 := 12
	myStr := "go is very fun"
	fmt.Println(booli,num,num2,myStr)

	for  i:= 0;i<10;i++{
		fmt.Println(i)
	}
	if x := 8;x%2 == 0{
		print("even")
	}
	val := 1
	switch val{
	case 1:
		print("one")
	case 2:
		print("two")
	default:
		print("other")
	}
}