package mylib

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello!")
}

func Sayhuman() {
	fmt.Println("human")
}
