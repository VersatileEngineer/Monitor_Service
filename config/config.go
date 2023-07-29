package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

// Load configuration from file
func LoadConfig(filepath string) error {
	Config = viper.New()
	Config.SetConfigFile(filepath)
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err.Error())
	}
	return nil
}

// Get blockchain type from configuration
func GetBlockchainType() string {
	return Config.GetString("blockchain.type")
}

// Get blockchain endpoint from configuration
func GetBlockchainEndpoint() string {
	return Config.GetString("blockchain.endpoint")
}

// Get database connection URL from configuration
func GetDatabaseConnection() string {
	return Config.GetString("database.url")
}

// Get message queue URL from configuration
func GetMessageQueueURL() string {
	return Config.GetString("message_queue.url")
}
