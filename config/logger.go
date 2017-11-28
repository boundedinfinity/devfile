package config

import (
	"log"
	"os"
)

func CreateLogger() *log.Logger {
	// prefix == "" and flag == 0, don't print any prefix information
	return log.New(os.Stdout, "", 0)
}
