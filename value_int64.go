package conf

// Int64Var defines an int64 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Int64Var(p *int64, name string, value int64, usage string) {
	c.env().Int64Var(p, name, value, usage)
	c.flag().Int64Var(p, name, value, usage)
}

// Int64 defines an int64 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)

	c.Int64Var(p, name, value, usage)

	return p
}

// Int64VarE defines an int64 environment variable with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the environment variable.
func (c *Configurator) Int64VarE(p *int64, name string, value int64, usage string) {
	c.env().Int64Var(p, name, value, usage)
}

// Int64E defines an int64 environment variable with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the environment variable.
func (c *Configurator) Int64E(name string, value int64, usage string) *int64 {
	p := new(int64)

	c.Int64VarE(p, name, value, usage)

	return p
}

// Int64VarF defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (c *Configurator) Int64VarF(p *int64, name string, value int64, usage string) {
	c.flag().Int64Var(p, name, value, usage)
}

// Int64F defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func (c *Configurator) Int64F(name string, value int64, usage string) *int64 {
	p := new(int64)

	c.Int64VarF(p, name, value, usage)

	return p
}

// Int64Var defines an int64 flag and environment variable with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag and/or environment variable.
func Int64Var(p *int64, name string, value int64, usage string) {
	Global.Int64Var(p, name, value, usage)
}

// Int64 defines an int64 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag and/or environment variable.
func Int64(name string, value int64, usage string) *int64 {
	return Global.Int64(name, value, usage)
}

// Int64VarE defines an int64 environment variable with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the environment variable.
func Int64VarE(p *int64, name string, value int64, usage string) {
	Global.Int64VarE(p, name, value, usage)
}

// Int64E defines an int64 environment variable with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the environment variable.
func Int64E(name string, value int64, usage string) *int64 {
	return Global.Int64E(name, value, usage)
}

// Int64VarF defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64VarF(p *int64, name string, value int64, usage string) {
	Global.Int64VarF(p, name, value, usage)
}

// Int64F defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64F(name string, value int64, usage string) *int64 {
	return Global.Int64F(name, value, usage)
}
