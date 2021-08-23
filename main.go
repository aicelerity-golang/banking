package main

import (
	"github.com/aicelerity-golang/banking/apps"
	"github.com/aicelerity-golang/banking/logger"
)

func main() {
	logger.Info("Starting our Go application")
	apps.Start()
}
