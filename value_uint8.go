package conf

// Uint8Var defines a uint8 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Uint8Var(p *uint8, name string, value uint8, usage string) {
	c.env().Uint8Var(p, name, value, usage)
	c.flag().Uint8Var(p, name, value, usage)
}

// Uint8 defines a uint8 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Uint8(name string, value uint8, usage string) *uint8 {
	p := new(uint8)

	c.Uint8Var(p, name, value, usage)

	return p
}

// Uint8VarE defines a uint8 environment variable with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the environment variable.
func (c *Configurator) Uint8VarE(p *uint8, name string, value uint8, usage string) {
	c.env().Uint8Var(p, name, value, usage)
}

// Uint8E defines a uint8 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the environment variable.
func (c *Configurator) Uint8E(name string, value uint8, usage string) *uint8 {
	p := new(uint8)

	c.Uint8VarE(p, name, value, usage)

	return p
}

// Uint8VarF defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func (c *Configurator) Uint8VarF(p *uint8, name string, value uint8, usage string) {
	c.flag().Uint8Var(p, name, value, usage)
}

// Uint8F defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func (c *Configurator) Uint8F(name string, value uint8, usage string) *uint8 {
	p := new(uint8)

	c.Uint8VarF(p, name, value, usage)

	return p
}

// Uint8Var defines a uint8 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag and/or environment variable.
func Uint8Var(p *uint8, name string, value uint8, usage string) {
	Global.Uint8Var(p, name, value, usage)
}

// Uint8 defines a uint8 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag and/or environment variable.
func Uint8(name string, value uint8, usage string) *uint8 {
	return Global.Uint8(name, value, usage)
}

// Uint8VarE defines a uint8 environment variable with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the environment variable.
func Uint8VarE(p *uint8, name string, value uint8, usage string) {
	Global.Uint8VarE(p, name, value, usage)
}

// Uint8E defines a uint8 environment variable with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the environment variable.
func Uint8E(name string, value uint8, usage string) *uint8 {
	return Global.Uint8E(name, value, usage)
}

// Uint8VarF defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func Uint8VarF(p *uint8, name string, value uint8, usage string) {
	Global.Uint8VarF(p, name, value, usage)
}

// Uint8F defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func Uint8F(name string, value uint8, usage string) *uint8 {
	return Global.Uint8F(name, value, usage)
}
