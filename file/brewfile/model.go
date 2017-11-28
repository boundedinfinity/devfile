package brewfile

import "encoding/xml"

type BrewfileLineType string

const (
    BrewfileLineType_Unknown BrewfileLineType = "unknown"
    BrewfileLineType_Brew    BrewfileLineType = "brew"
    BrewfileLineType_Cask    BrewfileLineType = "cask"
    BrewfileLineType_Tap     BrewfileLineType = "tap"
    BrewfileLineType_Comment BrewfileLineType = "#"
)

var (
    BrewfileLineTypes = []BrewfileLineType{
        BrewfileLineType_Brew,
        BrewfileLineType_Cask,
        BrewfileLineType_Tap,
        BrewfileLineType_Comment,
    }
)

func string2BrewfileLineType(s string) BrewfileLineType {
    t := BrewfileLineType_Unknown

    for _, c := range BrewfileLineTypes {
        if s == string(c) {
            t = c
            break
        }
    }

    return t
}

type BrewFileFile struct {
    XMLName  xml.Name       `json:"-" yaml:"-" xml:"devfile"`
    Path     string         `json:"path" yaml:"path" xml:"path"`
    Lines    []BrewFileLine `json:"lines" yaml:"lines" xml:"lines"`
    Actions  int            `json:"actions" yaml:"actions" xml:"actions"`
    Comments int            `json:"comments" yaml:"comments" xml:"comments"`
    Ignored  int            `json:"ignored" yaml:"ignored" xml:"ignored"`
}

func (this *BrewFileFile) AppendLine(t BrewfileLineType, v string, c string) {
    this.Lines = append(this.Lines, BrewFileLine{
        Name:    t,
        Value:   v,
        Comment: c,
    })
}

type BrewFileLine struct {
    Name    BrewfileLineType `json:"name" yaml:"name" xml:"name"`
    Value   string           `json:"value" yaml:"value" xml:"value"`
    Comment string           `json:"comment" yaml:"comment" xml:"comment"`
}
