package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/krislender0104/Monitor_Service/blockchain"
)

// Config represents the application configuration
type Config struct {
	BlockchainType     string                    `json:"blockchain_type"`
	BitcoinConfig      blockchain.BitcoinConfig  `json:"bitcoin_config"`
	EthereumConfig     blockchain.EthereumConfig `json:"ethereum_config"`
	RabbitMQURL        string                    `json:"rabbitmq_url"`
	SQLiteDatabasePath string                    `json:"sqlite_database_path"`
	Addresses          []string                  `json:"addresses"`
}

func loadConfigFromFile(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
