package conf

// Float32Var defines a float32 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) Float32Var(p *float32, name string, value float32, usage string) {
	c.env().Float32Var(p, name, value, usage)
	c.flag().Float32Var(p, name, value, usage)
}

// Float32 defines a float32 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Float32(name string, value float32, usage string) *float32 {
	p := new(float32)

	c.Float32Var(p, name, value, usage)

	return p
}

// Float32VarE defines a float32 environment variable with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the environment variable.
func (c *Configurator) Float32VarE(p *float32, name string, value float32, usage string) {
	c.env().Float32Var(p, name, value, usage)
}

// Float32E defines a float32 environment variable with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the environment variable.
func (c *Configurator) Float32E(name string, value float32, usage string) *float32 {
	p := new(float32)

	c.Float32VarE(p, name, value, usage)

	return p
}

// Float32VarF defines a float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag.
func (c *Configurator) Float32VarF(p *float32, name string, value float32, usage string) {
	c.flag().Float32Var(p, name, value, usage)
}

// Float32F defines a float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag.
func (c *Configurator) Float32F(name string, value float32, usage string) *float32 {
	p := new(float32)

	c.Float32VarF(p, name, value, usage)

	return p
}

// Float32Var defines a float32 flag and environment variable with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag and/or environment variable.
func Float32Var(p *float32, name string, value float32, usage string) {
	Global.Float32Var(p, name, value, usage)
}

// Float32 defines a float32 flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag and/or environment variable.
func Float32(name string, value float32, usage string) *float32 {
	return Global.Float32(name, value, usage)
}

// Float32VarE defines a float32 environment variable with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the environment variable.
func Float32VarE(p *float32, name string, value float32, usage string) {
	Global.Float32VarE(p, name, value, usage)
}

// Float32E defines a float32 environment variable with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the environment variable.
func Float32E(name string, value float32, usage string) *float32 {
	return Global.Float32E(name, value, usage)
}

// Float32VarF defines a float32 flag with specified name, default value, and usage string.
// The argument p points to a float32 variable in which to store the value of the flag.
func Float32VarF(p *float32, name string, value float32, usage string) {
	Global.Float32VarF(p, name, value, usage)
}

// Float32F defines a float32 flag with specified name, default value, and usage string.
// The return value is the address of a float32 variable that stores the value of the flag.
func Float32F(name string, value float32, usage string) *float32 {
	return Global.Float32F(name, value, usage)
}
