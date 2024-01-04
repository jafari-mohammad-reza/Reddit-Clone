package main

import (
	"github.com/reddit-clone/src"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)

func main() {
	cfg := config.GetConfig()
	lg := custome_logger.NewLogger(cfg)
	src.InitApp(cfg , lg)
	lg.Info(custome_logger.General, custome_logger.Startup, "Application started", nil)
}
