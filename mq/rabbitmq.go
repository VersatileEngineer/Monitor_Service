package mq

type RabbitMQ struct {
	URL string
	// Add necessary fields for managing RabbitMQ connection here
}

func NewRabbitMQ(url string) *RabbitMQ {
	return &RabbitMQ{
		URL: url,
		// Initialize RabbitMQ connection here
	}
}

func (q *RabbitMQ) Close() {
	// Close RabbitMQ connection here
}
