package conf

import "strings"

type queryStringValue map[string]string

func newQueryStringValue(val map[string]string, p *map[string]string) *queryStringValue {
	*p = val

	return (*queryStringValue)(p)
}

func (p queryStringValue) Set(val string) error {
	// Clear the map from default values
	for k := range p {
		delete(p, k)
	}

	for _, v := range strings.Split(val, "&") {
		param := strings.SplitN(v, "=", 2)
		if len(param) != 2 {
			continue
		}

		p[param[0]] = param[1]
	}

	return nil
}

func (queryStringValue) Type() string {
	return "queryString"
}

func (p queryStringValue) String() string {
	var query string

	for key, value := range p {
		if query != "" {
			query += "&"
		}

		query += key + "=" + value
	}

	return query
}

// QueryStringVar defines a query string flag and environment variable with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the flag and/or environment variable.
func (c *Configurator) QueryStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	c.env().QueryStringVar(p, name, value, usage)
	c.flag().Var(newQueryStringValue(value, p), name, usage)
}

// QueryString defines a query string flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the flag and/or environment variable.
func (c *Configurator) QueryString(name string, value map[string]string, usage string) *map[string]string {
	p := new(map[string]string)

	c.QueryStringVar(p, name, value, usage)

	return p
}

// QueryStringVarE defines a query string environment variable with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the environment variable.
func (c *Configurator) QueryStringVarE(p *map[string]string, name string, value map[string]string, usage string) {
	c.env().QueryStringVar(p, name, value, usage)
}

// QueryStringE defines a query string environment variable with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the environment variable.
func (c *Configurator) QueryStringE(name string, value map[string]string, usage string) *map[string]string {
	p := new(map[string]string)

	c.QueryStringVarE(p, name, value, usage)

	return p
}

// QueryStringVarF defines a query string flag with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the flag.
func (c *Configurator) QueryStringVarF(p *map[string]string, name string, value map[string]string, usage string) {
	c.flag().Var(newQueryStringValue(value, p), name, usage)
}

// QueryStringF defines a query string flag with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the flag.
func (c *Configurator) QueryStringF(name string, value map[string]string, usage string) *map[string]string {
	p := new(map[string]string)

	c.QueryStringVarF(p, name, value, usage)

	return p
}

// QueryStringVar defines a query string flag and environment variable with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the flag and/or environment variable.
func QueryStringVar(p *map[string]string, name string, value map[string]string, usage string) {
	Global.QueryStringVar(p, name, value, usage)
}

// QueryString defines a query string flag and environment variable with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the flag and/or environment variable.
func QueryString(name string, value map[string]string, usage string) *map[string]string {
	return Global.QueryString(name, value, usage)
}

// QueryStringVarE defines a query string environment variable with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the environment variable.
func QueryStringVarE(p *map[string]string, name string, value map[string]string, usage string) {
	Global.QueryStringVarE(p, name, value, usage)
}

// QueryStringE defines a query string environment variable with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the environment variable.
func QueryStringE(name string, value map[string]string, usage string) *map[string]string {
	return Global.QueryStringE(name, value, usage)
}

// QueryStringVarF defines a query string flag with specified name, default value, and usage string.
// The argument p points to a query string (string map) variable in which to store the value of the flag.
func QueryStringVarF(p *map[string]string, name string, value map[string]string, usage string) {
	Global.QueryStringVarF(p, name, value, usage)
}

// QueryStringF defines a query string flag with specified name, default value, and usage string.
// The return value is the address of a query string (string map) variable that stores the value of the flag.
func QueryStringF(name string, value map[string]string, usage string) *map[string]string {
	return Global.QueryStringF(name, value, usage)
}
