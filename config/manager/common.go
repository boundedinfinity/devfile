package manager

type stringConfigDescriptor struct {
	ShortFlag        string
	LongFlag         string
	Default          string
	ShortDescription string
	LongDescription  string
}

type boolConfigDescriptor struct {
	ShortFlag        string
	LongFlag         string
	Default          bool
	ShortDescription string
	LongDescription  string
}

type intConfigDescriptor struct {
	ShortFlag        string
	LongFlag         string
	Default          int
	ShortDescription string
	LongDescription  string
}

func (this *ConfigurationManager) configureBoolFlag(d boolConfigDescriptor) {
	if d.ShortFlag == "" {
		this.fs.Bool(d.LongFlag, d.Default, d.ShortDescription)
	} else {
		this.fs.BoolP(d.LongFlag, d.ShortFlag, d.Default, d.ShortDescription)
	}

	this.v.BindPFlag(d.LongFlag, this.fs.Lookup(d.LongFlag))
}

func (this *ConfigurationManager) configureStringFlag(d stringConfigDescriptor) {
	if d.ShortFlag == "" {
		this.fs.String(d.LongFlag, d.Default, d.ShortDescription)
	} else {
		this.fs.StringP(d.LongFlag, d.ShortFlag, d.Default, d.ShortDescription)
	}

	this.v.BindPFlag(d.LongFlag, this.fs.Lookup(d.LongFlag))
}

func (this *ConfigurationManager) configureIntFlag(d intConfigDescriptor) {
	if d.ShortFlag == "" {
		this.fs.Int(d.LongFlag, d.Default, d.ShortDescription)
	} else {
		this.fs.IntP(d.LongFlag, d.ShortFlag, d.Default, d.ShortDescription)
	}

	this.v.BindPFlag(d.LongFlag, this.fs.Lookup(d.LongFlag))
}
