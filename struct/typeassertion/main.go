package main

import (
	"fmt"
)

// 引数にいろいろな型を持たせたい時に,interface{}を引数に設定する。
// その後、type assertion
func do(i interface{}) {
	ii := i.(int) // type assertion. intであることを確認
	ii *= 2
	fmt.Println(ii)
}

// switch type
func do2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I do not know %T\n", v)
	}
}

func main() {
	do(10)
	do2(10)
	do2("Mike")
	do2(true)
}
