package main

import (
	"log"

	"github.com/krislender0104/Monitor_Service/blockchain"
	"github.com/krislender0104/Monitor_Service/config"
	"github.com/krislender0104/Monitor_Service/database"
	"github.com/krislender0104/Monitor_Service/messagequeue"
	"github.com/krislender0104/Monitor_Service/utils"
)

func main() {
	// Load configuration
	err := config.LoadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	// Initialize logger
	utils.InitLogger()

	// Create instances of blockchain, database, message queue, and backup services
	blockchainService := blockchain.NewBlockchainService()
	databaseService := database.NewDatabaseService()
	messagequeueService := messagequeue.NewMessageQueueService()
	backupService := backup.NewBackupService()

	// Connect to blockchain based on configuration
	err = blockchainService.ConnectToBlockchain(config.GetBlockchainType(), config.GetBlockchainEndpoint())
	if err != nil {
		log.Fatalf("Error connecting to blockchain: %s", err.Error())
	}

	// Connect to database
	err = databaseService.ConnectToDatabase(config.GetDatabaseConnection())
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	// Connect to message queue
	err = messagequeueService.ConnectToMessageQueue(config.GetMessageQueueURL())
	if err != nil {
		log.Fatalf("Error connecting to message queue: %s", err.Error())
	}

	// Start monitoring addresses
	messagequeueService.StartMonitoringAddresses(blockchainService, databaseService)

	// Backup in case of failure
	backupService.PerformBackup(databaseService)
}
