package conf

// Int8Var defines an int8 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Int8Var(p *int8, name string, value int8, usage string) {
	c.env().Int8Var(p, name, value, usage)
	c.flag().Int8Var(p, name, value, usage)
}

// Int8 defines an int8 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Int8(name string, value int8, usage string) *int8 {
	p := new(int8)

	c.Int8Var(p, name, value, usage)

	return p
}

// Int8VarE defines an int8 environment variable with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the environment variable.
func (c *Configurator) Int8VarE(p *int8, name string, value int8, usage string) {
	c.env().Int8Var(p, name, value, usage)
}

// Int8E defines an int8 environment variable with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the environment variable.
func (c *Configurator) Int8E(name string, value int8, usage string) *int8 {
	p := new(int8)

	c.Int8VarE(p, name, value, usage)

	return p
}

// Int8VarF defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func (c *Configurator) Int8VarF(p *int8, name string, value int8, usage string) {
	c.flag().Int8Var(p, name, value, usage)
}

// Int8F defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func (c *Configurator) Int8F(name string, value int8, usage string) *int8 {
	p := new(int8)

	c.Int8VarF(p, name, value, usage)

	return p
}

// Int8Var defines an int8 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag and/or environment variable.
func Int8Var(p *int8, name string, value int8, usage string) {
	Global.Int8Var(p, name, value, usage)
}

// Int8 defines an int8 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag and/or environment variable.
func Int8(name string, value int8, usage string) *int8 {
	return Global.Int8(name, value, usage)
}

// Int8VarE defines an int8 environment variable with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the environment variable.
func Int8VarE(p *int8, name string, value int8, usage string) {
	Global.Int8VarE(p, name, value, usage)
}

// Int8E defines an int8 environment variable with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the environment variable.
func Int8E(name string, value int8, usage string) *int8 {
	return Global.Int8E(name, value, usage)
}

// Int8VarF defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func Int8VarF(p *int8, name string, value int8, usage string) {
	Global.Int8VarF(p, name, value, usage)
}

// Int8F defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func Int8F(name string, value int8, usage string) *int8 {
	return Global.Int8F(name, value, usage)
}
