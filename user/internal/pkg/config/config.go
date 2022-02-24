package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type dbConf struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

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

	if viper.GetString("POSTGRES_DB_SERVER") == "" {
		err := loadDevEnv()
		if err != nil {
			return fmt.Errorf("Ошибка загрузки переменных окружения для разработки из %s : %w ", devEnvPath, err)
		}
	}

	return nil
}

func SetupPaths(cPath, cName, dPath string) {
	devEnvPath = dPath
	configPath = cPath
	configName = cName
}

func GetPostgreDSN() string {
	dbC := dbConf{
		Host:     viper.GetString("POSTGRES_DB_SERVER"),
		Port:     viper.GetString("POSTGRES_DB_SERVER_PORT"),
		Username: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		DbName:   viper.GetString("POSTGRES_DB"),
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbC.Username, dbC.Password, dbC.Host, dbC.Port, dbC.DbName)
}

func loadDevEnv() error {
	fmt.Println("load dev.env")
	return godotenv.Load(devEnvPath)
}
