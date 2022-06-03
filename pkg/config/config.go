package config

import (
	"awesomeProject/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Database struct {
		DbLogin    string `yaml:"db_login" env-default:"postgres"`
		DbPassword string `yaml:"db_password" env-default:"pavel"`
		DbHost     string `yaml:"db_host" env-default:"localhost"`
		DbPort     string `yaml:"db_port" env-default:"5432"`
		DbName     string `yaml:"db_name" env-default:"nft"`
	} `yaml:"database"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
