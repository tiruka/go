// fmt.Print系の出力内容を変更する方法
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	fmt.Println(mike)
}
