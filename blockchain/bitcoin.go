package blockchain

type BitcoinConfig struct {
	RegTestMode bool `json:"regtest_mode"`
}

type BitcoinClient struct {
	Config BitcoinConfig
}

func NewBitcoinClient(config BitcoinConfig) *BitcoinClient {
	return &BitcoinClient{
		Config: config,
	}
}

func (c *BitcoinClient) MonitorAddresses(addresses []string, rabbitMQ mq.MessageQueue, db database.Database) error {
	// Implementation for monitoring Bitcoin addresses
	return nil
}
