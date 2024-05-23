package main

import "fmt"


func sum(a int) func(int) int{
	return func(b int) int{
		return a * b
	}
}

func main(){
	fmt.Println("this is currying")
	a := sum(10)(12)
	fmt.Println(a)
}
