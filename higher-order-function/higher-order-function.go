package main

import "fmt"


func wrapperFunc(inner func()) func() {
	return func(){
		inner()
		fmt.Println("with wrapper")

	}
}

func main(){
	fmt.Println("this is hof")

	inner:=func(){ 
		fmt.Println("inner function")
	}

	wrapperFunc(inner)()
}
