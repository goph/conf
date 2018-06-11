package conf

// Uint64Var defines a uint64 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Uint64Var(p *uint64, name string, value uint64, usage string) {
	c.env().Uint64Var(p, name, value, usage)
	c.flag().Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)

	c.Uint64Var(p, name, value, usage)

	return p
}

// Uint64VarE defines a uint64 environment variable with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the environment variable.
func (c *Configurator) Uint64VarE(p *uint64, name string, value uint64, usage string) {
	c.env().Uint64Var(p, name, value, usage)
}

// Uint64E defines a uint64 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the environment variable.
func (c *Configurator) Uint64E(name string, value uint64, usage string) *uint64 {
	p := new(uint64)

	c.Uint64VarE(p, name, value, usage)

	return p
}

// Uint64VarF defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (c *Configurator) Uint64VarF(p *uint64, name string, value uint64, usage string) {
	c.flag().Uint64Var(p, name, value, usage)
}

// Uint64F defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (c *Configurator) Uint64F(name string, value uint64, usage string) *uint64 {
	p := new(uint64)

	c.Uint64VarF(p, name, value, usage)

	return p
}

// Uint64Var defines a uint64 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag and/or environment variable.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	Global.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag and/or environment variable.
func Uint64(name string, value uint64, usage string) *uint64 {
	return Global.Uint64(name, value, usage)
}

// Uint64VarE defines a uint64 environment variable with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the environment variable.
func Uint64VarE(p *uint64, name string, value uint64, usage string) {
	Global.Uint64VarE(p, name, value, usage)
}

// Uint64E defines a uint64 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the environment variable.
func Uint64E(name string, value uint64, usage string) *uint64 {
	return Global.Uint64E(name, value, usage)
}

// Uint64VarF defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64VarF(p *uint64, name string, value uint64, usage string) {
	Global.Uint64VarF(p, name, value, usage)
}

// Uint64F defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64F(name string, value uint64, usage string) *uint64 {
	return Global.Uint64F(name, value, usage)
}
