package config

type Config struct {
	BlockchainType string
	RabbitMQURL    string
	DatabaseURL    string
}

func LoadConfig() Config {
	// Load configuration from environment variables or config file
	// and return a Config struct
}
