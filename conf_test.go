package conf_test

import (
	"bytes"
	"testing"

	"github.com/goph/conf"
	flag "github.com/spf13/pflag"
)

type valueStub struct {
	err error
	typ string

	value string
}

func (v *valueStub) String() string {
	return v.value
}

func (v *valueStub) Set(value string) error {
	v.value = value

	return v.err
}

func (v *valueStub) Type() string {
	return v.typ
}

func TestFlagsTakePrecedence(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	c.Var(v, "value", "Value usage string")

	err := c.Parse(
		[]string{"name", "--value", "flag"},
		map[string]string{
			"VALUE": "env",
		},
	)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
	}

	if v.value != "flag" {
		t.Error("returned value is expected to be flag")
	}
}

func TestUsage(t *testing.T) {
	var buf bytes.Buffer

	c := conf.NewConfigurator("name", conf.ContinueOnError)
	c.SetOutput(&buf)

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	c.Var(v, "value", "Value usage string")

	err := c.Parse(
		[]string{"name", "--help"},
		map[string]string{},
	)

	if err != flag.ErrHelp {
		t.Fatal("Parse is expected to return a flag.ErrHelp when showing usage information")
	}

	expected := `Usage of name:

FLAGS

      --value valueStub   Value usage string


ENVIRONMENT VARIABLES

      VALUE valueStub   Value usage string
`

	actual := buf.String()

	if expected != actual {
		t.Errorf("expected the usage output: \n%s\n\nactual: \n%s", expected, actual)
	}
}

func TestConfigurator_LookupE(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	c.VarE(v, "value", "Value usage string")

	envVar := c.LookupE("value")

	if envVar == nil {
		t.Error("returned value is expected to be an EnVar, got nil")
	}

	if envVar.Value != v {
		t.Error("returned EnvVar is expected to be hold value")
	}
}

func TestConfigurator_LookupF(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)

	v := &valueStub{
		typ: "valueStub",
		err: nil,
	}

	c.VarF(v, "value", "Value usage string")

	flag := c.LookupF("value")

	if flag == nil {
		t.Error("returned value is expected to be an Flag, got nil")
	}

	if flag.Value != v {
		t.Error("returned Flag is expected to be hold value")
	}
}
