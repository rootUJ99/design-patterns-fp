package main

import "fmt"


func square(x int) int {
	return x * x
}

func plus5(x int) int {
	return x + 5
}

func compose(fn1 func(int) int, fn2 func(int) int) func(int)int{
	return func(y int) int{
		// res1 := fn1(y)
		// res2 := fn2(res1)
		// return res2
		return fn2(fn1(y))
	}
}

func main(){
	fmt.Println("this is composition")

	composedFunc := compose(square, plus5)
	res := composedFunc(10)
	fmt.Println(res)

}
