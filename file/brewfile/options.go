package brewfile

import (
	"log"
	"errors"
	"path/filepath"
    "github.com/boundedinfinity/devfile/config/manager"
)

type BrewFileProcessorOption func(*BrewFileProcessor) error

func Logger(input *log.Logger) BrewFileProcessorOption {
	return func(this *BrewFileProcessor) error {
		this.Logger = input
		return nil
	}
}

var PathIsEmptyError = errors.New("Path cannot be empty")

func Path(input string) BrewFileProcessorOption {
	return func(this *BrewFileProcessor) error {
		absPath, err := filepath.Abs(input)

		if err != nil {
			return err
		}

		this.Path = absPath
		return nil
	}
}

func Debug(input bool) BrewFileProcessorOption {
	return func(this *BrewFileProcessor) error {
		this.Debug = input
		return nil
	}
}

func OutputFormat(input manager.OutputFormat) BrewFileProcessorOption {
	return func(this *BrewFileProcessor) error {
		this.OutputFormat = input
		return nil
	}
}

func Clean(input bool) BrewFileProcessorOption {
	return func(this *BrewFileProcessor) error {
		this.Clean = input
		return nil
	}
}
