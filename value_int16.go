package conf

// Int16Var defines an int16 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Int16Var(p *int16, name string, value int16, usage string) {
	c.env().Int16Var(p, name, value, usage)
	c.flag().Int16Var(p, name, value, usage)
}

// Int16 defines an int16 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Int16(name string, value int16, usage string) *int16 {
	p := new(int16)

	c.Int16Var(p, name, value, usage)

	return p
}

// Int16VarE defines an int16 environment variable with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the environment variable.
func (c *Configurator) Int16VarE(p *int16, name string, value int16, usage string) {
	c.env().Int16Var(p, name, value, usage)
}

// Int16E defines an int16 environment variable with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the environment variable.
func (c *Configurator) Int16E(name string, value int16, usage string) *int16 {
	p := new(int16)

	c.Int16VarE(p, name, value, usage)

	return p
}

// Int16VarF defines an int16 flag with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag.
func (c *Configurator) Int16VarF(p *int16, name string, value int16, usage string) {
	c.flag().Int16Var(p, name, value, usage)
}

// Int16F defines an int16 flag with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the flag.
func (c *Configurator) Int16F(name string, value int16, usage string) *int16 {
	p := new(int16)

	c.Int16VarF(p, name, value, usage)

	return p
}

// Int16Var defines an int16 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag and/or environment variable.
func Int16Var(p *int16, name string, value int16, usage string) {
	Global.Int16Var(p, name, value, usage)
}

// Int16 defines an int16 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the flag and/or environment variable.
func Int16(name string, value int16, usage string) *int16 {
	return Global.Int16(name, value, usage)
}

// Int16VarE defines an int16 environment variable with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the environment variable.
func Int16VarE(p *int16, name string, value int16, usage string) {
	Global.Int16VarE(p, name, value, usage)
}

// Int16E defines an int16 environment variable with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the environment variable.
func Int16E(name string, value int16, usage string) *int16 {
	return Global.Int16E(name, value, usage)
}

// Int16VarF defines an int16 flag with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag.
func Int16VarF(p *int16, name string, value int16, usage string) {
	Global.Int16VarF(p, name, value, usage)
}

// Int16F defines an int16 flag with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the flag.
func Int16F(name string, value int16, usage string) *int16 {
	return Global.Int16F(name, value, usage)
}
