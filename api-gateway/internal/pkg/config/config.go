package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	devEnvPath = "./config/dev.env"
	configPath = "./config"
	configName = "config"
)

func Init() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Ошибка загрузки конфига: %w ", err)
	}

	return nil
}

func SetupPaths(cPath, cName, dPath string) {
	devEnvPath = dPath
	configPath = cPath
	configName = cName
}
