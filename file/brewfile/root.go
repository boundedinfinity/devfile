package brewfile

import (
    "log"
    "github.com/boundedinfinity/devfile/config"
    "github.com/boundedinfinity/devfile/config/manager"
)

type BrewFileProcessor struct {
    Path         string
    File         *BrewFileFile
    Debug        bool
    Logger       *log.Logger
    OutputFormat manager.OutputFormat
    Clean        bool
}

func NewBrewFileProcessor(options ...BrewFileProcessorOption) (*BrewFileProcessor, error) {
    service := &BrewFileProcessor{}

    for _, option := range options {
        if err := option(service); err != nil {
            return nil, err
        }
    }

    if service.Path == "" {
        return nil, PathIsEmptyError
    }

    if service.Logger == nil {
        service.Logger = config.CreateLogger()
    }

    service.File = &BrewFileFile{
        Path:  service.Path,
        Lines: make([]BrewFileLine, 0),
    }

    if service.Debug {
        service.Logger.Printf("Path: %s", service.Path)
        service.Logger.Printf("Debug: %t", service.Debug)
        service.Logger.Printf("Output Format: %t", service.OutputFormat)
        service.Logger.Printf("Clean: %t", service.Clean)
    }

    return service, nil
}

func (this *BrewFileProcessor) Execute() error {
    if err := this.lex(); err != nil {
        return err
    }

    if this.OutputFormat == manager.OutputFormat_None {
        this.Logger.Printf("Processed %d lines (%d actions, %d comments, %d ignored)",
            len(this.File.Lines), this.File.Actions, this.File.Comments, this.File.Ignored)
    } else {
        if err := manager.PrintFormatString(this.OutputFormat, this.File); err != nil {
            return err
        }
    }

    if this.Clean {
        if err := this.write(); err != nil {
            return err
        }
    }

    return nil
}
