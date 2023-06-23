package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	logger *Logger
)

// Its executed before the main func
func Init() error {
	err := loadConfig(".")
	if err != nil {
		return fmt.Errorf("error loading config: ", err)
	}
	return nil
}

func loadConfig(path string) error {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
