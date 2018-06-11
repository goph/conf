package conf

// Float64Var defines a float64 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Float64Var(p *float64, name string, value float64, usage string) {
	c.env().Float64Var(p, name, value, usage)
	c.flag().Float64Var(p, name, value, usage)
}

// Float64 defines a float64 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)

	c.Float64Var(p, name, value, usage)

	return p
}

// Float64VarE defines a float64 environment variable with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the environment variable.
func (c *Configurator) Float64VarE(p *float64, name string, value float64, usage string) {
	c.env().Float64Var(p, name, value, usage)
}

// Float64E defines a float64 environment variable with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the environment variable.
func (c *Configurator) Float64E(name string, value float64, usage string) *float64 {
	p := new(float64)

	c.Float64VarE(p, name, value, usage)

	return p
}

// Float64VarF defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (c *Configurator) Float64VarF(p *float64, name string, value float64, usage string) {
	c.flag().Float64Var(p, name, value, usage)
}

// Float64F defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (c *Configurator) Float64F(name string, value float64, usage string) *float64 {
	p := new(float64)

	c.Float64VarF(p, name, value, usage)

	return p
}

// Float64Var defines a float64 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag and/or environment variable.
func Float64Var(p *float64, name string, value float64, usage string) {
	Global.Float64Var(p, name, value, usage)
}

// Float64 defines a float64 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag and/or environment variable.
func Float64(name string, value float64, usage string) *float64 {
	return Global.Float64(name, value, usage)
}

// Float64VarE defines a float64 environment variable with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the environment variable.
func Float64VarE(p *float64, name string, value float64, usage string) {
	Global.Float64VarE(p, name, value, usage)
}

// Float64E defines a float64 environment variable with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the environment variable.
func Float64E(name string, value float64, usage string) *float64 {
	return Global.Float64E(name, value, usage)
}

// Float64VarF defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64VarF(p *float64, name string, value float64, usage string) {
	Global.Float64VarF(p, name, value, usage)
}

// Float64F defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64F(name string, value float64, usage string) *float64 {
	return Global.Float64F(name, value, usage)
}
