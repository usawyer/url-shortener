package util

import (
	"errors"
	"flag"
)

func ParseFlag() (string, error) {
	flagStore := flag.String("d", "", "Storage to store")
	flag.Parse()

	switch *flagStore {
	case "memory":
		return "memory", nil
	case "db":
		return "db", nil
	default:
		return "", errors.New("wrong flag")
	}
}
