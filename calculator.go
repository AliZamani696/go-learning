package main
import (
	"fmt"
	"math"
)



func calc( a float64,b float64, op string)float64{
	var res float64 =0
	switch  op {
	case "+":
		res = a + b
	case "/":
		if b ==0{
			fmt.Println("error zero division")
			return 0
		}
		res = a / b
	case "*":
		res = a * b
	case "**":
		res =math.Pow(a,b)
	}
	return res
}


func main(){
	fmt.Println(calc(2,3,"**"))
}