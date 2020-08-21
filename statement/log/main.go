package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// LoggingSetting file logging
func LoggingSetting(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

// ErrorLogging error
func ErrorLogging() {
	file, err := os.Open("not_exist_file")
	if err != nil {
		log.Fatalln("Exit with error", err)
	}
	fmt.Fprintln(file)
}

func main() {
	// ErrorLogging()
	LoggingSetting("test.log")
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")
	log.Fatalf("%T %v", "test", "test") //Fatal系を使うと、ここで終了する。
	log.Fatalln("error")
	log.Println("ok") // not print
}
