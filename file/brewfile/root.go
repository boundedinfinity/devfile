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

    return service, nil
}

func (this *BrewFileProcessor) Read() error {
    if this.Debug {
        this.Logger.Printf("Path: %s", this.Path)
        this.Logger.Printf("Debug: %t", this.Debug)
        this.Logger.Printf("Output Format: %t", this.OutputFormat)
    }

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

    return nil
}

func (this *BrewFileProcessor) Write(l BrewFileLine) error {
    if err := this.lex(); err != nil {
        return err
    }

    if err := this.File.AppendLine(l); err != nil {
        return err
    }

    if err := this.write(); err != nil {
        return err
    }

    return nil
}
