package main

import (
	"Reddit-Clone/src"
	"Reddit-Clone/src/share/config"
	"Reddit-Clone/src/share/pkg/custome_logger"
)

func main() {
	cfg := config.GetConfig()
	lg := custome_logger.NewLogger(cfg)
	src.InitApp(cfg, lg)
	lg.Info(custome_logger.General, custome_logger.Startup, "Application started", nil)
}
