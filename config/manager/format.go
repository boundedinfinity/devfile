package manager

import (
    "strings"
    "fmt"
    "encoding/xml"
    "encoding/json"
    "github.com/pkg/errors"
    "gopkg.in/yaml.v2"
)

type OutputFormat string

const (
    OutputFormat_None OutputFormat = ""
    OutputFormat_Json OutputFormat = "json"
    OutputFormat_Yaml OutputFormat = "yaml"
    OutputFormat_Xml  OutputFormat = "xml"
)

var (
    OutputFormats = []OutputFormat{ OutputFormat_Json, OutputFormat_Yaml, OutputFormat_Xml, }

	jsonDescriptor = stringConfigDescriptor{
		ShortFlag:        "f",
		LongFlag:         "format",
		ShortDescription: fmt.Sprintf("command output format.  Can be one of [%s]", strings.Join(outputFormats2Strings(), ",")),
		LongDescription:  fmt.Sprintf("command output format.  Can be one of [%s]", strings.Join(outputFormats2Strings(), ",")),
	}
)

func outputFormats2Strings() []string {
    ofs := make([]string, 0)
    for _, of := range OutputFormats {
        ofs = append(ofs, string(of))
    }
    return ofs
}

func (this *ConfigurationManager) ConfigureFormat() {
	this.configureStringFlag(jsonDescriptor)
}

func (this *ConfigurationManager) GetFormat() OutputFormat {
    raw := this.v.GetString(jsonDescriptor.LongFlag)
    of := OutputFormat_None

    for _, cof := range OutputFormats {
        if raw == string(cof) {
            of = cof
            break
        }
    }
    return of
}

func (this *ConfigurationManager) ToFormatString(v interface{}) (string, error) {
    return ToFormatString(this.GetFormat(), v)
}

func (this *ConfigurationManager) PrintFormatString(v interface{}) error {
    return PrintFormatString(this.GetFormat(), v)
}

func ToFormatString(of OutputFormat, v interface{}) (string, error) {
    if of == OutputFormat_Yaml {
        b, err := yaml.Marshal(v);
        return string(b), err
    }

    if of == OutputFormat_Xml {
        b, err := xml.Marshal(v);
        return string(b), err
    }

    if of == OutputFormat_Json {
        b, err := json.Marshal(v);
        return string(b), err
    }

    return "", errors.Errorf("invalid format: %v", of)
}

func PrintFormatString(of OutputFormat, v interface{}) error {
    s, err := ToFormatString(of, v)

    if err != nil {
        return err
    }

    fmt.Print(s)
    return nil
}

func (this *ConfigurationManager) ValidateFormat() error {
    raw := this.v.GetString(jsonDescriptor.LongFlag)
    of := OutputFormat_None

    for _, cof := range OutputFormats {
        if raw == string(cof) {
            of = cof
            break
        }
    }

    if of == OutputFormat_None {
        return errors.Errorf("invalid output format: %s, must be one of [%s]",
            raw, strings.Join(outputFormats2Strings(), ","))
    }

    return nil
}
