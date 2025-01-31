package main

import (
	"github.com/VitorSantosCruz/gopportunities/config"
	"github.com/VitorSantosCruz/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
