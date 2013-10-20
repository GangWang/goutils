package main

import (
	"logger"
)

func main(){
	logger.Init(4,"log")
	logger.Info("info")
	logger.Debug("debug")
	logger.Warn("warn")
	logger.Error("error")
} 