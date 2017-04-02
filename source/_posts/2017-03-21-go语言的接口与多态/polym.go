package main

import "fmt"

type Animal interface {
	eat()
}

type Tiger struct {
	food string
}

func (t Tiger) eat() {
	fmt.Println("tiger eat ", t.food)
}
type Monkey struct {
	food string
}
func (m Monkey) eat() {
	fmt.Println("monkey eat ", m.food)
}

func main() {
	var animal Animal
	tiger := Tiger{"meat"}
	monkey := Monkey{"banana"}
	animal = tiger
	animal.eat()
	animal = monkey
	animal.eat()
}
