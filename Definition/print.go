package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("bazz")
}

func main() {
	bazz()
	fmt.Println("Hello world with go", time.Now())
	fmt.Println(user.Current())
}
