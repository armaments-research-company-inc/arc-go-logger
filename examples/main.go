package main

import (
	logger "github.com/armaments-research-company-inc/arc-go-logger"
)

func main() {

	logger.Initialize(nil)
	logger.Info("Hello!")
}
