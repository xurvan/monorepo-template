package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port        uint16
	DatabaseUrl string
}

func LoadConfig() (cfg *Config) {
	v := viper.New()
	v.AutomaticEnv()

	port := v.GetUint16("PORT")
	databaseUrl := v.GetString("DATABASE_URL")

	cfg = &Config{
		Port:        port,
		DatabaseUrl: databaseUrl,
	}

	return cfg
}
