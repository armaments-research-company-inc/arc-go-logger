package main

import (
	logger "gitlab.com/arc-rte/arc-go-logger"
)

func main() {

	logger.Initialize(nil)
	logger.Info("Hello!")
}
