package utils

import (
	"log"
)

// Handle errors and log them
func HandleError(err error) {
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
}
