package backup

type RedisClient struct{}

func (redis *RedisClient) Backup() {
	// Backup microservice state to Redis in case of failure
}
