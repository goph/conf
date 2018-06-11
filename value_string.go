package conf

// StringVar defines a string flag and environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) StringVar(p *string, name string, value string, usage string) {
	c.env().StringVar(p, name, value, usage)
	c.flag().StringVar(p, name, value, usage)
}

// String defines a string flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag and/or environment variable.
func (c *Configurator) String(name string, value string, usage string) *string {
	p := new(string)

	c.StringVar(p, name, value, usage)

	return p
}

// StringVarE defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the environment variable.
func (c *Configurator) StringVarE(p *string, name string, value string, usage string) {
	c.env().StringVar(p, name, value, usage)
}

// StringE defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the environment variable.
func (c *Configurator) StringE(name string, value string, usage string) *string {
	p := new(string)

	c.StringVarE(p, name, value, usage)

	return p
}

// StringVarF defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (c *Configurator) StringVarF(p *string, name string, value string, usage string) {
	c.flag().StringVar(p, name, value, usage)
}

// StringF defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (c *Configurator) StringF(name string, value string, usage string) *string {
	p := new(string)

	c.StringVarF(p, name, value, usage)

	return p
}

// StringVar defines a string flag and environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag and/or environment variable.
func StringVar(p *string, name string, value string, usage string) {
	Global.StringVar(p, name, value, usage)
}

// String defines a string flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag and/or environment variable.
func String(name string, value string, usage string) *string {
	return Global.String(name, value, usage)
}

// StringVarE defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the environment variable.
func StringVarE(p *string, name string, value string, usage string) {
	Global.StringVarE(p, name, value, usage)
}

// StringE defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the environment variable.
func StringE(name string, value string, usage string) *string {
	return Global.StringE(name, value, usage)
}

// StringVarF defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVarF(p *string, name string, value string, usage string) {
	Global.StringVarF(p, name, value, usage)
}

// StringF defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func StringF(name string, value string, usage string) *string {
	return Global.StringF(name, value, usage)
}
