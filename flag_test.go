package conf_test

import (
	"testing"

	"time"

	"github.com/goph/conf"
)

func TestFlagValue(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)
	vars := new(valueVars)

	vars.boolVar = c.BoolF("bool", false, "bool value")
	vars.durationVar = c.DurationF("duration", 0, "time.Duration value")
	vars.float32Var = c.Float32F("float32", 0, "float32 value")
	vars.float64Var = c.Float64F("float64", 0, "float64 value")
	vars.intVar = c.IntF("int", 0, "int value")
	vars.int16Var = c.Int16F("int16", 0, "int16 value")
	vars.int32Var = c.Int32F("int32", 0, "int32 value")
	vars.int64Var = c.Int64F("int64", 0, "int64 value")
	vars.int8Var = c.Int8F("int8", 0, "int8 value")
	vars.stringVar = c.StringF("string", "", "string value")
	vars.uintVar = c.UintF("uint", 0, "uint value")
	vars.uint16Var = c.Uint16F("uint16", 0, "uint16 value")
	vars.uint32Var = c.Uint32F("uint32", 0, "uint32 value")
	vars.uint64Var = c.Uint64F("uint64", 0, "uint64 value")
	vars.uint8Var = c.Uint8F("uint8", 0, "uint8 value")

	testFlagValue(t, c, vars)
}

func TestFlagValueVar(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)
	vars := newValueVars()

	c.BoolVarF(vars.boolVar, "bool", false, "bool value")
	c.DurationVarF(vars.durationVar, "duration", 0, "time.Duration value")
	c.Float32VarF(vars.float32Var, "float32", 0, "float32 value")
	c.Float64VarF(vars.float64Var, "float64", 0, "float64 value")
	c.IntVarF(vars.intVar, "int", 0, "int value")
	c.Int16VarF(vars.int16Var, "int16", 0, "int16 value")
	c.Int32VarF(vars.int32Var, "int32", 0, "int32 value")
	c.Int64VarF(vars.int64Var, "int64", 0, "int64 value")
	c.Int8VarF(vars.int8Var, "int8", 0, "int8 value")
	c.StringVarF(vars.stringVar, "string", "", "string value")
	c.UintVarF(vars.uintVar, "uint", 0, "uint value")
	c.Uint16VarF(vars.uint16Var, "uint16", 0, "uint16 value")
	c.Uint32VarF(vars.uint32Var, "uint32", 0, "uint32 value")
	c.Uint64VarF(vars.uint64Var, "uint64", 0, "uint64 value")
	c.Uint8VarF(vars.uint8Var, "uint8", 0, "uint8 value")

	testFlagValue(t, c, vars)
}

func TestGlobalFlagValue(t *testing.T) {
	conf.Global = conf.NewConfigurator("name", conf.ContinueOnError)
	vars := new(valueVars)

	vars.boolVar = conf.BoolF("bool", false, "bool value")
	vars.durationVar = conf.DurationF("duration", 0, "time.Duration value")
	vars.float32Var = conf.Float32F("float32", 0, "float32 value")
	vars.float64Var = conf.Float64F("float64", 0, "float64 value")
	vars.intVar = conf.IntF("int", 0, "int value")
	vars.int16Var = conf.Int16F("int16", 0, "int16 value")
	vars.int32Var = conf.Int32F("int32", 0, "int32 value")
	vars.int64Var = conf.Int64F("int64", 0, "int64 value")
	vars.int8Var = conf.Int8F("int8", 0, "int8 value")
	vars.stringVar = conf.StringF("string", "", "string value")
	vars.uintVar = conf.UintF("uint", 0, "uint value")
	vars.uint16Var = conf.Uint16F("uint16", 0, "uint16 value")
	vars.uint32Var = conf.Uint32F("uint32", 0, "uint32 value")
	vars.uint64Var = conf.Uint64F("uint64", 0, "uint64 value")
	vars.uint8Var = conf.Uint8F("uint8", 0, "uint8 value")

	testFlagValue(t, conf.Global, vars)
}

func TestGlobalFlagValueVar(t *testing.T) {
	conf.Global = conf.NewConfigurator("name", conf.ContinueOnError)
	vars := newValueVars()

	conf.BoolVarF(vars.boolVar, "bool", false, "bool value")
	conf.DurationVarF(vars.durationVar, "duration", 0, "time.Duration value")
	conf.Float32VarF(vars.float32Var, "float32", 0, "float32 value")
	conf.Float64VarF(vars.float64Var, "float64", 0, "float64 value")
	conf.IntVarF(vars.intVar, "int", 0, "int value")
	conf.Int16VarF(vars.int16Var, "int16", 0, "int16 value")
	conf.Int32VarF(vars.int32Var, "int32", 0, "int32 value")
	conf.Int64VarF(vars.int64Var, "int64", 0, "int64 value")
	conf.Int8VarF(vars.int8Var, "int8", 0, "int8 value")
	conf.StringVarF(vars.stringVar, "string", "", "string value")
	conf.UintVarF(vars.uintVar, "uint", 0, "uint value")
	conf.Uint16VarF(vars.uint16Var, "uint16", 0, "uint16 value")
	conf.Uint32VarF(vars.uint32Var, "uint32", 0, "uint32 value")
	conf.Uint64VarF(vars.uint64Var, "uint64", 0, "uint64 value")
	conf.Uint8VarF(vars.uint8Var, "uint8", 0, "uint8 value")

	testFlagValue(t, conf.Global, vars)
}

func testFlagValue(t *testing.T, c *conf.Configurator, vars *valueVars) {
	arguments := []string{
		"name",
		"--bool", "true",
		"--duration", "1s",
		"--float32", "172e12",
		"--float64", "2718e28",
		"--int", "22",
		"--int16", "16",
		"--int32", "32",
		"--int64", "64",
		"--int8", "8",
		"--string", "string",
		"--uint", "22",
		"--uint16", "16",
		"--uint32", "32",
		"--uint64", "64",
		"--uint8", "8",
	}

	environment := map[string]string{}

	err := c.Parse(arguments, environment)

	if err != nil {
		t.Fatalf("Parse is expected to return a nil (non-error) value, got: %s", err)
	}

	if !*vars.boolVar {
		t.Error("bool var should be `true`, got: ", *vars.boolVar)
	}

	if *vars.durationVar != time.Second {
		t.Error("duration var should be `1s`, got: ", (*vars.durationVar).String())
	}

	if *vars.float32Var != 172e12 {
		t.Error("float32 var should be `172e12`, got: ", *vars.float32Var)
	}

	if *vars.float64Var != 2718e28 {
		t.Error("float64 var should be `172e12`, got: ", *vars.float64Var)
	}

	if *vars.intVar != 22 {
		t.Error("int var should be `22`, got: ", *vars.intVar)
	}

	if *vars.int16Var != 16 {
		t.Error("int16 var should be `16`, got: ", *vars.int16Var)
	}

	if *vars.int32Var != 32 {
		t.Error("int32 var should be `32`, got: ", *vars.int32Var)
	}

	if *vars.int64Var != 64 {
		t.Error("int64 var should be `64`, got: ", *vars.int64Var)
	}

	if *vars.int8Var != 8 {
		t.Error("int8 var should be `8`, got: ", *vars.int8Var)
	}

	if *vars.stringVar != "string" {
		t.Error("string var should be `string`, got: ", *vars.stringVar)
	}

	if *vars.uintVar != 22 {
		t.Error("uint var should be `22`, got: ", *vars.uintVar)
	}

	if *vars.uint16Var != 16 {
		t.Error("uint16 var should be `16`, got: ", *vars.uint16Var)
	}

	if *vars.uint32Var != 32 {
		t.Error("uint32 var should be `32`, got: ", *vars.uint32Var)
	}

	if *vars.uint64Var != 64 {
		t.Error("uint64 var should be `64`, got: ", *vars.uint64Var)
	}

	if *vars.uint8Var != 8 {
		t.Error("uint8 var should be `8`, got: ", *vars.uint8Var)
	}
}
