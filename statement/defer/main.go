package main

import (
	"fmt"
	"os"
)

func some() {
	defer fmt.Println("Hello World") // after completed this func, deal this defer
	fmt.Println("Hello Go")
}

// deferはファイルをopen, closeするときなどに便利下記のように使う
func filehandler() {
	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

func main() {
	// defer some()
	// fmt.Println("Main func")
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3) // stacking process
	fmt.Println("success")
	filehandler()
}
