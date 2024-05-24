package main

import (
	"fmt"
	"slices"
)

type Trail struct {
	value int
	logs []string
}

func initiateTrail(value int) *Trail{
	return &Trail{
		value: value,
		logs: []string{},
	}	
}

func square(value int) Trail {
	return Trail{
		value: value * value,
		logs: []string{"value has been squared"},
	}
}

func addFive(value int) Trail {
	return Trail{
		value: value + 5,
		logs: []string{"five has been added to value"},
	}
}

func (t *Trail) transform(fn func(int)Trail) *Trail {
	transformedValue := fn(t.value)
	return &Trail{
		value: transformedValue.value,
		logs: slices.Concat(transformedValue.logs, t.logs),
	}
}

func main(){
	fmt.Println("this is monad")

	trail := initiateTrail(5).transform(square).transform(addFive)
	fmt.Println(trail)
}
