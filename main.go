package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/krislender0104/Monitor_Service/blockchain"
	"github.com/krislender0104/Monitor_Service/database"
	"github.com/krislender0104/Monitor_Service/mq"
)

func main() {
	//Initialize configuration
	config, err := loadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	//Create connections to RabbitMQ and SQLite
	rabbitMQ := mq.NewRabbitMQ(config.RabbitMQURL)
	sqliteDB := database.NewSQLiteDB(config.SQLiteDatabasePath)

	//Initialize blockchain clients based on configuration
	var blockchainClient blockchain.BlockChain
	switch config.BlockchainType {
	case "bitcoin":
		blockchainClient = blockchain.NewBitcoinClient(config.BitcoinConfig)
	case "ethereum":
		blockchainClient = blockchain.NewEthereumClient(config.EthereumConfig)
	default:
		log.Fatalf("Invalid blockchain type: %s", config.BlockchainType)
	}

	// Start address monitoring
	err = blockchainClient.MonitorAddresses(config.Addresses, rabbitMQ, sqliteDB)
	if err != nil {
		log.Fatalf("Failed to start address monitoring: %s", err)
	}

	//Wait for termination siganl to gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	//Cleanup resources
	rabbitMQ.Close()
	sqliteDB.Close()
}

func loadConfig() (*config, error) {
	//Load configuration from file or environment variables
	return &config{
		BlockchainType:     "bitcoin",
		BitcoinConfig:      blockchain.BitcoinConfig{RegTestMode: true},
		EthereumConfig:     blockchain.EthereumConfig{Testnet: true},
		RabbitMQURL:        "amqp://guest:guest@localhost:5672/",
		SQLiteDatabasePath: "data.db",
		Addresses:          []string{""},
	}, nil
}

type config struct {
	BlockchainType     string
	BitcoinConfig      blockchain.BitcoinConfig
	EthereumConfig     blockchain.EthereumConfig
	RabbitMQURL        string
	SQLiteDatabasePath string
	Addresses          []string
}
