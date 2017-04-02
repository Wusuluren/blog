package main

import "fmt"

type ArryElem interface {
	dummy()
}

type IntElem int
func (i IntElem) dummy() {
}

type FloatElem float64
func (f FloatElem) dummy() {
}

func main() {
	arry := make([]ArryElem, 3)
	for i := 0; i < 3; i++ {
		arry[i] = IntElem(i)
	}
	for i := range arry {
		fmt.Println(arry[i])
	}
	for i := 0; i < 3; i++ {
		arry[i] = FloatElem(i)+0.1
	}
	for i := range arry {
		fmt.Println(arry[i])
	}
}