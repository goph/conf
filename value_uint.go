package conf

// UintVar defines a uint flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) UintVar(p *uint, name string, value uint, usage string) {
	c.env().UintVar(p, name, value, usage)
	c.flag().UintVar(p, name, value, usage)
}

// Uint defines a uint flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Uint(name string, value uint, usage string) *uint {
	p := new(uint)

	c.UintVar(p, name, value, usage)

	return p
}

// UintVarE defines a uint environment variable with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the environment variable.
func (c *Configurator) UintVarE(p *uint, name string, value uint, usage string) {
	c.env().UintVar(p, name, value, usage)
}

// UintE defines a uint environment variable with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the environment variable.
func (c *Configurator) UintE(name string, value uint, usage string) *uint {
	p := new(uint)

	c.UintVarE(p, name, value, usage)

	return p
}

// UintVarF defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (c *Configurator) UintVarF(p *uint, name string, value uint, usage string) {
	c.flag().UintVar(p, name, value, usage)
}

// UintF defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func (c *Configurator) UintF(name string, value uint, usage string) *uint {
	p := new(uint)

	c.UintVarF(p, name, value, usage)

	return p
}

// UintVar defines a uint flag and environment variable with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag and/or environment variable.
func UintVar(p *uint, name string, value uint, usage string) {
	Global.UintVar(p, name, value, usage)
}

// Uint defines a uint flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag and/or environment variable.
func Uint(name string, value uint, usage string) *uint {
	return Global.Uint(name, value, usage)
}

// UintVarE defines a uint environment variable with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the environment variable.
func UintVarE(p *uint, name string, value uint, usage string) {
	Global.UintVarE(p, name, value, usage)
}

// UintE defines a uint environment variable with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the environment variable.
func UintE(name string, value uint, usage string) *uint {
	return Global.UintE(name, value, usage)
}

// UintVarF defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func UintVarF(p *uint, name string, value uint, usage string) {
	Global.UintVarF(p, name, value, usage)
}

// UintF defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func UintF(name string, value uint, usage string) *uint {
	return Global.UintF(name, value, usage)
}
