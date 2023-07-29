package blockchain

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

type BitcoinClient struct{}

func ConnectToRegTest() {
	// Connect to Bitcoin RegTest mode
	connConfig := &rpcclient.ConnConfig{
		Host:         "localhost:18443", // Replace with your regtest node's address
		User:         "yourusername",    // Replace with your regtest node's username
		Pass:         "yourpassword",    // Replace with your regtest node's password
		HTTPPostMode: true,
		DisableTLS:   true, // Disable TLS if not using HTTPS
	}
	client, err := rpcclient.New(connConfig, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Make an API call to get the Bitcoin price
	price, err := client.GetDifficulty()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current difficulty: %v\n", price)
}
