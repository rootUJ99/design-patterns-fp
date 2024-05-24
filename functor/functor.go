package main

import "fmt"

type Binder interface {
	mapp()
}

type Option struct{
	value string
}

func(o Option) mapp(fn func(string)string) Option{
	return Option{value: fn(o.value)}
}

func newOption(str string) Option {
	return Option{
		value: str, 
	}
}

func main(){
	fmt.Println("this is functor")

	abc:= newOption("abc")
	abc2:=abc.mapp(func(str string)string{
		return fmt.Sprintf("%s 1", str)
	})
	fmt.Println(abc, abc2)
}
