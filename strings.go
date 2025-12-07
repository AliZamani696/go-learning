package main

import (
	"fmt"
	"strings"
	// "unicode/utf8"
)



func main(){
	s := "hello"
	s2 := "world"

	s3 := "hello go syntax"

	s += ", " + s2

	f := fmt.Sprintf(s,s2)


	words := []int{1,24}
	fmt.Println(f,words)
	str := "hello mello"

	charLen := utf8.RuneCountInString(str)
	text := "[import] message [len]"
	start := strings.Index(text,"[")+1
	end := strings.Index(text,"]")
	email := "azamani081@gmail.com"
	dom :=strings.Split(email,"@")[0]
	ho := strings.Split(email,"@")[1]
	fmt.Println(dom,ho)

	splitFunc := func(c rune) bool{
		return c == ',' || c == ';'
	}
	fmt.Println(strings.FieldsFunc("a,dfg;d,d",splitFunc))

	var buider strings.Builder
	buider.Grow(100)
	for i := 0;i<100;i++{
		if i % 2 ==0{
			buider.WriteString("*")
		}

	}
	res := buider.String()
	fmt.Println(res)

}