package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config holds the configuration settings for the application.
type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

// LoadConfig loads the configuration from the groceries.toml file.
func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("groceries") // name of config file (without extension)
	viper.SetConfigType("toml")       // required
	viper.AddConfigPath(".")          // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return config, err
	}

	return config, nil
}
