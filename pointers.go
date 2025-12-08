


package main

import "fmt"

func incptr(num *int64)(res bool){
	res = true
	for i :=int64(2);i<int64((*num/2)+1);i++{
		if *num % i ==0{
			res = false
			break
		}

	}
	return
}
func main(){
	 var i int64 =1
	for i < 10000000{
		 if incptr(&i){
			fmt.Println(i)
		 }
		i++
	}
}