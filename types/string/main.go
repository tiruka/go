package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("Hello" + "Go!")
	fmt.Println("Hello Go!"[0]) // ascii
	fmt.Println(string("Hello Go!"[0]))
	var s string = "Hello World"
	// s[0] = x  Error
	var y string = strings.Replace(s, "H", "X", 1)
	fmt.Println(y)
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(strings.Contains(s, "Go"))
	fmt.Println("Test" +
		"test")
	fmt.Println(`Test
		test`)
}
