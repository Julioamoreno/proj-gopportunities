package main

import (
	"github.com/Julioamoreno/proj-gopportunities/config"
	"github.com/Julioamoreno/proj-gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.ErrorF("Config initialization error: %v", err)

		return
	}

	router.Initialize()
}
