package conf

// BoolVar defines a bool flag and environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) BoolVar(p *bool, name string, value bool, usage string) {
	c.env().BoolVar(p, name, value, usage)
	c.flag().BoolVar(p, name, value, usage)
}

// Bool defines a bool flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Bool(name string, value bool, usage string) *bool {
	p := new(bool)

	c.BoolVar(p, name, value, usage)

	return p
}

// BoolVarE defines a bool environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the environment variable.
func (c *Configurator) BoolVarE(p *bool, name string, value bool, usage string) {
	c.env().BoolVar(p, name, value, usage)
}

// BoolE defines a bool environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the environment variable.
func (c *Configurator) BoolE(name string, value bool, usage string) *bool {
	p := new(bool)

	c.BoolVarE(p, name, value, usage)

	return p
}

// BoolVarF defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (c *Configurator) BoolVarF(p *bool, name string, value bool, usage string) {
	c.flag().BoolVar(p, name, value, usage)
}

// BoolF defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (c *Configurator) BoolF(name string, value bool, usage string) *bool {
	p := new(bool)

	c.BoolVarF(p, name, value, usage)

	return p
}

// BoolVar defines a bool flag and environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag and/or environment variable.
func BoolVar(p *bool, name string, value bool, usage string) {
	Global.BoolVar(p, name, value, usage)
}

// Bool defines a bool flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag and/or environment variable.
func Bool(name string, value bool, usage string) *bool {
	return Global.Bool(name, value, usage)
}

// BoolVarE defines a bool environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the environment variable.
func BoolVarE(p *bool, name string, value bool, usage string) {
	Global.BoolVarE(p, name, value, usage)
}

// BoolE defines a bool environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the environment variable.
func BoolE(name string, value bool, usage string) *bool {
	return Global.BoolE(name, value, usage)
}

// BoolVarF defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVarF(p *bool, name string, value bool, usage string) {
	Global.BoolVarF(p, name, value, usage)
}

// BoolF defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func BoolF(name string, value bool, usage string) *bool {
	return Global.BoolF(name, value, usage)
}
