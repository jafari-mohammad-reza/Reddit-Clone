package main

import (
	"Reddit-Clone/src"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//cfg := config.GetConfig()
	PORT := os.Args[1]
	port, err := strconv.Atoi(PORT)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(PORT)
	fmt.Println(port)
	//lg := custome_logger.NewLogger()
	src.InitApp(port)
	//lg.Info(custome_logger.General, custome_logger.Startup, "Application started", nil)
}
