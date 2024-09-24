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
