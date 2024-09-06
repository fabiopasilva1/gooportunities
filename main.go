package main

import (
	"github.com/fabiopasilva1/gooportunities/config"
	"github.com/fabiopasilva1/gooportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	// Initialize Router
	router.Initialize()
}
