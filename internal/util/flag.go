package util

import (
	"errors"
	"os"
)

func ParseFlag() (string, error) {
	flagStore := os.Getenv("STORAGE_FLAG")

	switch flagStore {
	case "memory":
		return "memory", nil
	case "db":
		return "db", nil
	default:
		return "", errors.New("wrong flag")
	}
}
