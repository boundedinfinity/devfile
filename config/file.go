package config

import (
	"os"
)

func GetFileMode() os.FileMode {
	return os.FileMode(0644)
}
