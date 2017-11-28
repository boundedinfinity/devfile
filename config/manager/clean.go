package manager

var (
	cleanDescriptor = boolConfigDescriptor{
		ShortFlag:        "c",
		LongFlag:         "clean",
		ShortDescription: "clean file contents",
		LongDescription:  "clean file contents",
		Default:          false,
	}
)

func (this *ConfigurationManager) ConfigureClean() {
	this.configureBoolFlag(cleanDescriptor)
}

func (this *ConfigurationManager) GetClean() bool {
	return this.v.GetBool(cleanDescriptor.LongFlag)
}
