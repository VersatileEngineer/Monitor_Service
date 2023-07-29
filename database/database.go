package database

type DatabaseClient struct{}

func (db *DatabaseClient) StoreData(data interface{}) {
	// Store monitored data in SQLite database
}
