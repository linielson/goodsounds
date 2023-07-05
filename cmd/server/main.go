package main

import (
	"github.com/linielson/goodsounds/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	logger.Infof("test: %v", config.GetTokenAuth())
	logger.Infof("test2: %v", config.GetJwtExpiresIn())
}
