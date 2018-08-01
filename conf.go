// Package conf provides easy CLI configuration management via flags and environment variables.
package conf

import (
	"fmt"
	"io"
	"os"

	"github.com/goph/env"
	flag "github.com/spf13/pflag"
)

// ErrorHandling defines how to handle parsing errors.
type ErrorHandling int

const (
	// ContinueOnError will return an err from Parse() if an error is found
	ContinueOnError ErrorHandling = iota

	// ExitOnError will call os.Exit(2) if an error is found
	ExitOnError

	// PanicOnError will panic() if an error is found
	PanicOnError
)

// Configurator exposes an interface to configure variables via flags and/or environment variables.
type Configurator struct {
	envVarSet *env.EnvVarSet // use env() accessor
	flagSet   *flag.FlagSet  // use flag() accessor

	name          string
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use out() accessor
}

// NewConfigurator returns a new configurator instance.
func NewConfigurator(name string, errorHandling ErrorHandling) *Configurator {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)

	flagSet.Usage = func() {}

	return &Configurator{
		envVarSet: env.NewEnvVarSet(env.ContinueOnError),
		flagSet:   flagSet,

		name:          name,
		errorHandling: errorHandling,
	}
}

func (c *Configurator) env() *env.EnvVarSet {
	if c.envVarSet == nil {
		c.envVarSet = env.NewEnvVarSet(env.ContinueOnError)
	}

	return c.envVarSet
}

func (c *Configurator) flag() *flag.FlagSet {
	if c.flagSet == nil {
		c.flagSet = flag.NewFlagSet("", flag.ContinueOnError)

		c.flagSet.Usage = func() {}
	}

	return c.flagSet
}

// Var defines a flag and an environment variable with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
func (c *Configurator) Var(value Value, name string, usage string) {
	c.env().Var(value, name, usage)
	c.flag().Var(value, name, usage)
}

// VarE defines an environment variable with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
func (c *Configurator) VarE(value Value, name string, usage string) {
	c.env().Var(value, name, usage)
}

// VarF defines a flag with the specified name and usage string.
// The type and value of the variable are represented by the first argument,
// of type Value, which typically holds a user-defined implementation of Value.
func (c *Configurator) VarF(value Value, name string, usage string) {
	c.flag().Var(value, name, usage)
}

func (c *Configurator) out() io.Writer {
	if c.output == nil {
		return os.Stderr
	}

	return c.output
}

// SetOutput sets the destination for usage and error messages in the underlying FlagSet and EnvVarSet.
func (c *Configurator) SetOutput(output io.Writer) {
	c.output = output

	c.env().SetOutput(output)
	c.flag().SetOutput(output)
}

// LookupE returns the EnvVar structure of the named environment variable, returning nil if none exists.
func (c *Configurator) LookupE(name string) *env.EnvVar {
	return c.env().Lookup(name)
}

// LookupF returns the Flag structure of the named flag, returning nil if none exists.
func (c *Configurator) LookupF(name string) *flag.Flag {
	return c.flag().Lookup(name)
}

// Parse parses flags and environment variables according to the definitions in the FlagSet and EnvVarSet.
// Must be called after all variables in the FlagSet and EnvVarSet
// are defined and before variables are accessed by the program.
func (c *Configurator) Parse(arguments []string, environment map[string]string) error {
	err := c.env().Parse(environment)
	if err != nil {
		switch c.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			fmt.Fprintln(c.out(), err)
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}

	err = c.flag().Parse(arguments)
	if err == flag.ErrHelp {
		c.usage()

		switch c.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError, PanicOnError:
			// Exit here without outputting the error
			os.Exit(2)
		}
	}

	if err != nil {
		switch c.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			fmt.Fprintln(c.out(), err)
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}

	return nil
}

// Parsed reports whether Parse has been called.
func (c *Configurator) Parsed() bool {
	return c.env().Parsed() && c.flag().Parsed()
}

func (c *Configurator) usage() {
	fmt.Fprintf(c.out(), "Usage of %s:\n\n", c.name)
	fmt.Fprintf(c.out(), "FLAGS\n\n%s\n\n", c.flag().FlagUsages())
	fmt.Fprintf(c.out(), "ENVIRONMENT VARIABLES\n\n%s", c.env().EnvVarUsages())
}
