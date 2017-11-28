package docker_compose

import (
    "encoding/xml"
    "bytes"
)

type DockerComposeVersion string

const (
    DockerComposeVersion_Unknown DockerComposeVersion = "-"
    DockerComposeVersion_3       DockerComposeVersion = "3"
)

var (
    DockerComposeVersions = []DockerComposeVersion{
        DockerComposeVersion_3,
    }
)

type DockerComposeFile struct {
    XMLName  xml.Name               `json:"-" yaml:"-" xml:"devfile"`
    Version  DockerComposeVersion   `json:"version" yaml:"version" xml:"version"`
    Services []DockerComposeService `json:"version" yaml:"version" xml:"version"`
}

type DockerComposeService struct {
    ServiceName  string
}

func (this *DockerComposeService) MarshalJSON() ([]byte, error) {
    buffer := bytes.NewBufferString("{\"")
    buffer.WriteString(this.ServiceName)
    buffer.WriteString("\"}")
    return buffer.Bytes(), nil
}
