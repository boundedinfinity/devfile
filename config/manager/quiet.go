package manager

var (
	quietDescriptor = boolConfigDescriptor{
		ShortFlag:        "q",
		LongFlag:         "quietDescriptor",
		ShortDescription: "enable quietDescriptor output",
		LongDescription:  "enable quietDescriptor output",
		Default:          false,
	}
)

func (this *ConfigurationManager) ConfigureQuiet() {
	this.configureBoolFlag(quietDescriptor)
}

func (this *ConfigurationManager) GetQuiet() bool {
	return this.v.GetBool(quietDescriptor.LongFlag)
}
