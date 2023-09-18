package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug bool          `yaml:"is_debug" env-required:"true"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() *Config {
	var instance *Config
	var once sync.Once

	once.Do(func() {

		instance = &Config{}

		if err := cleanenv.ReadConfig("src/internal/config/config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Printf("%s\n", err)
			log.Printf("%s\n", help)

		}
	})
	return instance
}
