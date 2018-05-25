package config

import (
	"github.com/gobuffalo/packr"
)

func Packr() packr.Box {
	return packr.NewBox("../templates")
}
