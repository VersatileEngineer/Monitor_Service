package main

import (
	"log"

	"github.com/krislender0104/Monitor_Service/blockchain"
	"github.com/krislender0104/Monitor_Service/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(cfg.BlockchainType)

	// Connect to blockchain based on configuration
	var client interface{}
	if cfg.BlockchainType == "bitcoin" {
		btcClient := blockchain.BitcoinClient{}
		btcClient.ConnectToRegTest()
		client = &btcClient
	} else {
		ethClient := blockchain.EthereumClient{}
		ethClient.ConnectToTestnet()
		client = &ethClient
	}

	// Initialize RabbitMQ client
	// rmqClient := rabbitmq.RabbitMQClient{}

	// Receive designated addresses from RabbitMQ and start monitoring transactions
	// go rmqClient.ReceiveAddresses()

	// Handle monitored data and store it in the database
	// dbClient := database.DatabaseClient{}
	// for {
	// 	// Get monitored data from blockchain client

	// 	// Store monitored data in database
	// 	dbClient.StoreData(data)

	// 	// Handle errors
	// 	errorhandling.HandleError(err)
	// }

	// // Log messages for debugging and issue resolution
	// logger := logging.Logger{}
	// logger.Log(message)

	// // Backup microservice state to Redis in case of failure
	// redisClient := backup.RedisClient{}
	// redisClient.Backup()
}
