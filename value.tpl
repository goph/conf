package conf

// {{Type}}Var defines a {{type}} flag and environment variable with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) {{Type}}Var(p *{{type}}, name string, value {{type}}, usage string) {
	c.env().{{Type}}Var(p, name, value, usage)
	c.flag().{{Type}}Var(p, name, value, usage)
}

// {{Type}} defines a {{type}} flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the flag and/or environment variable.
func (c *Configurator) {{Type}}(name string, value {{type}}, usage string) *{{type}} {
	p := new({{type}})

	c.{{Type}}Var(p, name, value, usage)

	return p
}

// {{Type}}VarE defines a {{type}} environment variable with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the environment variable.
func (c *Configurator) {{Type}}VarE(p *{{type}}, name string, value {{type}}, usage string) {
	c.env().{{Type}}Var(p, name, value, usage)
}

// {{Type}}E defines a {{type}} environment variable with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the environment variable.
func (c *Configurator) {{Type}}E(name string, value {{type}}, usage string) *{{type}} {
	p := new({{type}})

	c.{{Type}}VarE(p, name, value, usage)

	return p
}

// {{Type}}VarF defines a {{type}} flag with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the flag.
func (c *Configurator) {{Type}}VarF(p *{{type}}, name string, value {{type}}, usage string) {
	c.flag().{{Type}}Var(p, name, value, usage)
}

// {{Type}}F defines a {{type}} flag with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the flag.
func (c *Configurator) {{Type}}F(name string, value {{type}}, usage string) *{{type}} {
	p := new({{type}})

	c.{{Type}}VarF(p, name, value, usage)

	return p
}

// {{Type}}Var defines a {{type}} flag and environment variable with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the flag and/or environment variable.
func {{Type}}Var(p *{{type}}, name string, value {{type}}, usage string) {
	Global.{{Type}}Var(p, name, value, usage)
}

// {{Type}} defines a {{type}} flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the flag and/or environment variable.
func {{Type}}(name string, value {{type}}, usage string) *{{type}} {
	return Global.{{Type}}(name, value, usage)
}

// {{Type}}VarE defines a {{type}} environment variable with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the environment variable.
func {{Type}}VarE(p *{{type}}, name string, value {{type}}, usage string) {
	Global.{{Type}}VarE(p, name, value, usage)
}

// {{Type}}E defines a {{type}} environment variable with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the environment variable.
func {{Type}}E(name string, value {{type}}, usage string) *{{type}} {
	return Global.{{Type}}E(name, value, usage)
}

// {{Type}}VarF defines a {{type}} flag with specified name, default value, and usage string.
// The argument p points to a {{type}} variable in which to store the value of the flag.
func {{Type}}VarF(p *{{type}}, name string, value {{type}}, usage string) {
	Global.{{Type}}VarF(p, name, value, usage)
}

// {{Type}}F defines a {{type}} flag with specified name, default value, and usage string.
// The return value is the address of a {{type}} variable that stores the value of the flag.
func {{Type}}F(name string, value {{type}}, usage string) *{{type}} {
	return Global.{{Type}}F(name, value, usage)
}
