package manager

var (
	debugDescriptor = boolConfigDescriptor{
		ShortFlag:        "d",
		LongFlag:         "debug",
		ShortDescription: "enable debugging output",
		LongDescription:  "enable debugging output",
		Default:          false,
	}
)

func (this *ConfigurationManager) ConfigureDebug() {
	this.configureBoolFlag(debugDescriptor)
}

func (this *ConfigurationManager) GetDebug() bool {
	return this.v.GetBool(debugDescriptor.LongFlag)
}
