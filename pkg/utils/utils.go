package utils

import (
	"path/filepath"

	"github.com/google/uuid"
)

var (
	ROOT_DIR = filepath.Dir("../../..")
)

func GetStringUUIDv7() string {
	id, err := uuid.NewV7()
	if err != nil {
		Logger.Fatal("", err)
	}
	return id.String()
}

// Helper function to check if all elements in a slice are equal
func AllEqual[T comparable](slice []T) bool {
	if len(slice) == 0 {
		return true
	}
	first := slice[0]
	for _, item := range slice[1:] {
		if item != first {
			return false
		}
	}
	return true
}
