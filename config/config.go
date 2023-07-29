package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	BlockchainType string
	RabbitMQURL    string
	DatabaseURL    string
}

func LoadConfig() (Config, error) {
	viper.SetConfigFile("config.yaml") // Set the name of the config file (you can change it to your preferred format)

	// Add any necessary configuration defaults here
	viper.SetDefault("BlockchainType", "ethereum")
	viper.SetDefault("RabbitMQURL", "amqp://guest:guest@localhost:5672/")
	viper.SetDefault("DatabaseURL", "sqlite:///path/to/database.db")

	if err := viper.ReadInConfig(); err != nil { // Read the config file and handle errors
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	err := viper.Unmarshal(&cfg) // Unmarshal the config file into a Config struct
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return cfg, nil
}
