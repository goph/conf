package conf

// StringSliceVar defines a string flag and environment variable with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag and/or environment variable.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (c *Configurator) StringSliceVar(p *[]string, name string, value []string, usage string) {
	c.env().StringSliceVar(p, name, value, usage)
	c.flag().StringSliceVar(p, name, value, usage)
}

// StringSlice defines a string flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag and/or environment variable.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (c *Configurator) StringSlice(name string, value []string, usage string) *[]string {
	p := new([]string)

	c.StringSliceVar(p, name, value, usage)

	return p
}

// StringSliceVarE defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (c *Configurator) StringSliceVarE(p *[]string, name string, value []string, usage string) {
	c.env().StringSliceVar(p, name, value, usage)
}

// StringSliceE defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (c *Configurator) StringSliceE(name string, value []string, usage string) *[]string {
	p := new([]string)

	c.StringSliceVarE(p, name, value, usage)

	return p
}

// StringSliceVarF defines a string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (c *Configurator) StringSliceVarF(p *[]string, name string, value []string, usage string) {
	c.flag().StringSliceVar(p, name, value, usage)
}

// StringSliceF defines a string flag with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
// will result in
//   []string{"v1", "v2", "v3"}
func (c *Configurator) StringSliceF(name string, value []string, usage string) *[]string {
	p := new([]string)

	c.StringSliceVarF(p, name, value, usage)

	return p
}

// StringSliceVar defines a string flag and environment variable with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag and/or environment variable.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func StringSliceVar(p *[]string, name string, value []string, usage string) {
	Global.StringSliceVar(p, name, value, usage)
}

// StringSlice defines a string flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag and/or environment variable.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func StringSlice(name string, value []string, usage string) *[]string {
	return Global.StringSlice(name, value, usage)
}

// StringSliceVarE defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func StringSliceVarE(p *[]string, name string, value []string, usage string) {
	Global.StringSliceVarE(p, name, value, usage)
}

// StringSliceE defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the environment variable.
// For example:
//   STRING_SLICE="v1,v2,v3"
// will result in
//   []string{"v1", "v2", "v3"}
func StringSliceE(name string, value []string, usage string) *[]string {
	return Global.StringSliceE(name, value, usage)
}

// StringSliceVarF defines a string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
// will result in
//   []string{"v1", "v2", "v3"}
func StringSliceVarF(p *[]string, name string, value []string, usage string) {
	Global.StringSliceVarF(p, name, value, usage)
}

// StringSliceF defines a string flag with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag.
// For example:
//   --string-slice="v1,v2" --string-slice="v3"
// will result in
//   []string{"v1", "v2", "v3"}
func StringSliceF(name string, value []string, usage string) *[]string {
	return Global.StringSliceF(name, value, usage)
}
