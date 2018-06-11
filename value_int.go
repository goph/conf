package conf

// IntVar defines an int flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) IntVar(p *int, name string, value int, usage string) {
	c.env().IntVar(p, name, value, usage)
	c.flag().IntVar(p, name, value, usage)
}

// Int defines an int flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Int(name string, value int, usage string) *int {
	p := new(int)

	c.IntVar(p, name, value, usage)

	return p
}

// IntVarE defines an int environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the environment variable.
func (c *Configurator) IntVarE(p *int, name string, value int, usage string) {
	c.env().IntVar(p, name, value, usage)
}

// IntE defines an int environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the environment variable.
func (c *Configurator) IntE(name string, value int, usage string) *int {
	p := new(int)

	c.IntVarE(p, name, value, usage)

	return p
}

// IntVarF defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (c *Configurator) IntVarF(p *int, name string, value int, usage string) {
	c.flag().IntVar(p, name, value, usage)
}

// IntF defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (c *Configurator) IntF(name string, value int, usage string) *int {
	p := new(int)

	c.IntVarF(p, name, value, usage)

	return p
}

// IntVar defines an int flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag and/or environment variable.
func IntVar(p *int, name string, value int, usage string) {
	Global.IntVar(p, name, value, usage)
}

// Int defines an int flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag and/or environment variable.
func Int(name string, value int, usage string) *int {
	return Global.Int(name, value, usage)
}

// IntVarE defines an int environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the environment variable.
func IntVarE(p *int, name string, value int, usage string) {
	Global.IntVarE(p, name, value, usage)
}

// IntE defines an int environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the environment variable.
func IntE(name string, value int, usage string) *int {
	return Global.IntE(name, value, usage)
}

// IntVarF defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVarF(p *int, name string, value int, usage string) {
	Global.IntVarF(p, name, value, usage)
}

// IntF defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func IntF(name string, value int, usage string) *int {
	return Global.IntF(name, value, usage)
}
