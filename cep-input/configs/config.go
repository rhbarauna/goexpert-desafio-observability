package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	WEATHER_API_URL string `mapstructure:"WEATHER_API_URL"`
	TRACER_URL      string `mapstructure:"TRACER_URL"`
	SERVICE_NAME    string `mapstructure:"SERVICE_NAME"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}

	return cfg, err
}
