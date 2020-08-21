package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error for file")
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error for test")
	}
}
