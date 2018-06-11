package conf

// Int32Var defines an int32 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Int32Var(p *int32, name string, value int32, usage string) {
	c.env().Int32Var(p, name, value, usage)
	c.flag().Int32Var(p, name, value, usage)
}

// Int32 defines an int32 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Int32(name string, value int32, usage string) *int32 {
	p := new(int32)

	c.Int32Var(p, name, value, usage)

	return p
}

// Int32VarE defines an int32 environment variable with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the environment variable.
func (c *Configurator) Int32VarE(p *int32, name string, value int32, usage string) {
	c.env().Int32Var(p, name, value, usage)
}

// Int32E defines an int32 environment variable with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the environment variable.
func (c *Configurator) Int32E(name string, value int32, usage string) *int32 {
	p := new(int32)

	c.Int32VarE(p, name, value, usage)

	return p
}

// Int32VarF defines an int32 flag with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag.
func (c *Configurator) Int32VarF(p *int32, name string, value int32, usage string) {
	c.flag().Int32Var(p, name, value, usage)
}

// Int32F defines an int32 flag with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag.
func (c *Configurator) Int32F(name string, value int32, usage string) *int32 {
	p := new(int32)

	c.Int32VarF(p, name, value, usage)

	return p
}

// Int32Var defines an int32 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag and/or environment variable.
func Int32Var(p *int32, name string, value int32, usage string) {
	Global.Int32Var(p, name, value, usage)
}

// Int32 defines an int32 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag and/or environment variable.
func Int32(name string, value int32, usage string) *int32 {
	return Global.Int32(name, value, usage)
}

// Int32VarE defines an int32 environment variable with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the environment variable.
func Int32VarE(p *int32, name string, value int32, usage string) {
	Global.Int32VarE(p, name, value, usage)
}

// Int32E defines an int32 environment variable with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the environment variable.
func Int32E(name string, value int32, usage string) *int32 {
	return Global.Int32E(name, value, usage)
}

// Int32VarF defines an int32 flag with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag.
func Int32VarF(p *int32, name string, value int32, usage string) {
	Global.Int32VarF(p, name, value, usage)
}

// Int32F defines an int32 flag with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag.
func Int32F(name string, value int32, usage string) *int32 {
	return Global.Int32F(name, value, usage)
}
