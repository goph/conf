package conf

import "time"

// DurationVar defines a time.Duration flag and environment variable with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	c.env().DurationVar(p, name, value, usage)
	c.flag().DurationVar(p, name, value, usage)
}

// Duration defines a time.Duration flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag and/or environment variable.
func (c *Configurator) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)

	c.DurationVar(p, name, value, usage)

	return p
}

// DurationVarE defines a time.Duration environment variable with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the environment variable.
func (c *Configurator) DurationVarE(p *time.Duration, name string, value time.Duration, usage string) {
	c.env().DurationVar(p, name, value, usage)
}

// DurationE defines a time.Duration environment variable with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the environment variable.
func (c *Configurator) DurationE(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)

	c.DurationVarE(p, name, value, usage)

	return p
}

// DurationVarF defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (c *Configurator) DurationVarF(p *time.Duration, name string, value time.Duration, usage string) {
	c.flag().DurationVar(p, name, value, usage)
}

// DurationF defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (c *Configurator) DurationF(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)

	c.DurationVarF(p, name, value, usage)

	return p
}

// DurationVar defines a time.Duration flag and environment variable with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag and/or environment variable.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	Global.DurationVar(p, name, value, usage)
}

// Duration defines a time.Duration flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag and/or environment variable.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return Global.Duration(name, value, usage)
}

// DurationVarE defines a time.Duration environment variable with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the environment variable.
func DurationVarE(p *time.Duration, name string, value time.Duration, usage string) {
	Global.DurationVarE(p, name, value, usage)
}

// DurationE defines a time.Duration environment variable with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the environment variable.
func DurationE(name string, value time.Duration, usage string) *time.Duration {
	return Global.DurationE(name, value, usage)
}

// DurationVarF defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func DurationVarF(p *time.Duration, name string, value time.Duration, usage string) {
	Global.DurationVarF(p, name, value, usage)
}

// DurationF defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func DurationF(name string, value time.Duration, usage string) *time.Duration {
	return Global.DurationF(name, value, usage)
}
