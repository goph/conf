package conf

// Uint16Var defines a uint16 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Uint16Var(p *uint16, name string, value uint16, usage string) {
	c.env().Uint16Var(p, name, value, usage)
	c.flag().Uint16Var(p, name, value, usage)
}

// Uint16 defines a uint16 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Uint16(name string, value uint16, usage string) *uint16 {
	p := new(uint16)

	c.Uint16Var(p, name, value, usage)

	return p
}

// Uint16VarE defines a uint16 environment variable with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the environment variable.
func (c *Configurator) Uint16VarE(p *uint16, name string, value uint16, usage string) {
	c.env().Uint16Var(p, name, value, usage)
}

// Uint16E defines a uint16 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the environment variable.
func (c *Configurator) Uint16E(name string, value uint16, usage string) *uint16 {
	p := new(uint16)

	c.Uint16VarE(p, name, value, usage)

	return p
}

// Uint16VarF defines a uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag.
func (c *Configurator) Uint16VarF(p *uint16, name string, value uint16, usage string) {
	c.flag().Uint16Var(p, name, value, usage)
}

// Uint16F defines a uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag.
func (c *Configurator) Uint16F(name string, value uint16, usage string) *uint16 {
	p := new(uint16)

	c.Uint16VarF(p, name, value, usage)

	return p
}

// Uint16Var defines a uint16 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag and/or environment variable.
func Uint16Var(p *uint16, name string, value uint16, usage string) {
	Global.Uint16Var(p, name, value, usage)
}

// Uint16 defines a uint16 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag and/or environment variable.
func Uint16(name string, value uint16, usage string) *uint16 {
	return Global.Uint16(name, value, usage)
}

// Uint16VarE defines a uint16 environment variable with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the environment variable.
func Uint16VarE(p *uint16, name string, value uint16, usage string) {
	Global.Uint16VarE(p, name, value, usage)
}

// Uint16E defines a uint16 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the environment variable.
func Uint16E(name string, value uint16, usage string) *uint16 {
	return Global.Uint16E(name, value, usage)
}

// Uint16VarF defines a uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag.
func Uint16VarF(p *uint16, name string, value uint16, usage string) {
	Global.Uint16VarF(p, name, value, usage)
}

// Uint16F defines a uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag.
func Uint16F(name string, value uint16, usage string) *uint16 {
	return Global.Uint16F(name, value, usage)
}
