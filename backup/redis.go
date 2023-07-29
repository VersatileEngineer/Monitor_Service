package backup

import (
	"log"
)

type BackupService struct{}

// Perform backup in Redis
func (r *BackupService) PerformBackup(databaseService *DatabaseService) {
	// Implementation for performing backup in Redis
	log.Println("Performed backup in Redis")
}
