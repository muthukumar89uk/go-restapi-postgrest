package config

import (
	"github.com/spf13/viper"
)

func Configuration(configPath string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
