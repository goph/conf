package conf

// Uint32Var defines a uint32 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Uint32Var(p *uint32, name string, value uint32, usage string) {
	c.env().Uint32Var(p, name, value, usage)
	c.flag().Uint32Var(p, name, value, usage)
}

// Uint32 defines a uint32 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Uint32(name string, value uint32, usage string) *uint32 {
	p := new(uint32)

	c.Uint32Var(p, name, value, usage)

	return p
}

// Uint32VarE defines a uint32 environment variable with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the environment variable.
func (c *Configurator) Uint32VarE(p *uint32, name string, value uint32, usage string) {
	c.env().Uint32Var(p, name, value, usage)
}

// Uint32E defines a uint32 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the environment variable.
func (c *Configurator) Uint32E(name string, value uint32, usage string) *uint32 {
	p := new(uint32)

	c.Uint32VarE(p, name, value, usage)

	return p
}

// Uint32VarF defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func (c *Configurator) Uint32VarF(p *uint32, name string, value uint32, usage string) {
	c.flag().Uint32Var(p, name, value, usage)
}

// Uint32F defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag.
func (c *Configurator) Uint32F(name string, value uint32, usage string) *uint32 {
	p := new(uint32)

	c.Uint32VarF(p, name, value, usage)

	return p
}

// Uint32Var defines a uint32 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag and/or environment variable.
func Uint32Var(p *uint32, name string, value uint32, usage string) {
	Global.Uint32Var(p, name, value, usage)
}

// Uint32 defines a uint32 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag and/or environment variable.
func Uint32(name string, value uint32, usage string) *uint32 {
	return Global.Uint32(name, value, usage)
}

// Uint32VarE defines a uint32 environment variable with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the environment variable.
func Uint32VarE(p *uint32, name string, value uint32, usage string) {
	Global.Uint32VarE(p, name, value, usage)
}

// Uint32E defines a uint32 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the environment variable.
func Uint32E(name string, value uint32, usage string) *uint32 {
	return Global.Uint32E(name, value, usage)
}

// Uint32VarF defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func Uint32VarF(p *uint32, name string, value uint32, usage string) {
	Global.Uint32VarF(p, name, value, usage)
}

// Uint32F defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag.
func Uint32F(name string, value uint32, usage string) *uint32 {
	return Global.Uint32F(name, value, usage)
}
