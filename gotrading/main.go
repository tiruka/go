package main

import (
	"fmt"
	"go/gotrading/config"
	"go/gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
