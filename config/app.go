package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	Port            string
	AppName         string         `mapstructure:"app_name"`
	LogLevel        string         `mapstructure:"log_level"`
	LogPath         string         `mapstructure:"log_path"`
	SessionLifetime int            `mapstructure:"session_lifetime"`
	Database        DatabaseConfig `mapstructure:",squash"`
}

type DatabaseConfig struct {
	Name     string `mapstructure:"database_name"`
	Username string `mapstructure:"database_username"`
	Password string `mapstructure:"database_password"`
	Host     string `mapstructure:"database_host"`
}

func Load() (AppConfig, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading configuration file : %s\n", err)
		return AppConfig{}, err
	}

	config := AppConfig{}
	viper.Unmarshal(&config)
	return config, nil
}
