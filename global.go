package conf

import (
	"os"
	"strings"
)

// Global is the default set of flags and environment variables
var Global = NewConfigurator(os.Args[0], ExitOnError)

// Var defines a flag and an environment variable with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
func Var(value Value, name string, usage string) {
	Global.Var(value, name, usage)
}

// VarE defines an environment variable with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
func VarE(value Value, name string, usage string) {
	Global.VarE(value, name, usage)
}

// VarF defines a flag with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
func VarF(value Value, name string, usage string) {
	Global.VarF(value, name, usage)
}

// Parse parses flags and environment variables according to the definitions in the FlagSet and EnvVarSet.
// Must be called after all variables in the FlagSet and EnvVarSet
// are defined and before variables are accessed by the program.
func Parse() {
	environment := map[string]string{}
	for _, environ := range os.Environ() {
		e := strings.SplitN(environ, "=", 2)

		environment[e[0]] = e[1]
	}

	Global.Parse(os.Args[1:], environment) // nolint:errcheck
}

// Parsed reports whether Parse has been called.
func Parsed() bool {
	return Global.Parsed()
}
