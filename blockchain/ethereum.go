package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthereumClient struct{}

func (eth *EthereumClient) ConnectToTestnet() {
	// Connect to Ethereum Testnet or Ganache
	client, err := ethclient.Dial("http://localhost:8545") // Update with your Ganache URL
	if err != nil {
		log.Fatal(err)
	}

}
