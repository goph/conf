package conf_test

import (
	"testing"

	"time"

	"github.com/goph/conf"
)

func TestEnvValue(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)
	vars := new(valueVars)

	vars.boolVar = c.BoolE("bool", false, "bool value")
	vars.durationVar = c.DurationE("duration", 0, "time.Duration value")
	vars.float32Var = c.Float32E("float32", 0, "float32 value")
	vars.float64Var = c.Float64E("float64", 0, "float64 value")
	vars.intVar = c.IntE("int", 0, "int value")
	vars.int16Var = c.Int16E("int16", 0, "int16 value")
	vars.int32Var = c.Int32E("int32", 0, "int32 value")
	vars.int64Var = c.Int64E("int64", 0, "int64 value")
	vars.int8Var = c.Int8E("int8", 0, "int8 value")
	vars.queryStringVar = c.QueryStringE("query-string", map[string]string{}, "query string value")
	vars.stringVar = c.StringE("string", "", "string value")
	vars.uintVar = c.UintE("uint", 0, "uint value")
	vars.uint16Var = c.Uint16E("uint16", 0, "uint16 value")
	vars.uint32Var = c.Uint32E("uint32", 0, "uint32 value")
	vars.uint64Var = c.Uint64E("uint64", 0, "uint64 value")
	vars.uint8Var = c.Uint8E("uint8", 0, "uint8 value")

	testEnvValue(t, c, vars)
}

func TestEnvValueVar(t *testing.T) {
	c := conf.NewConfigurator("name", conf.ContinueOnError)
	vars := newValueVars()

	c.BoolVarE(vars.boolVar, "bool", false, "bool value")
	c.DurationVarE(vars.durationVar, "duration", 0, "time.Duration value")
	c.Float32VarE(vars.float32Var, "float32", 0, "float32 value")
	c.Float64VarE(vars.float64Var, "float64", 0, "float64 value")
	c.IntVarE(vars.intVar, "int", 0, "int value")
	c.Int16VarE(vars.int16Var, "int16", 0, "int16 value")
	c.Int32VarE(vars.int32Var, "int32", 0, "int32 value")
	c.Int64VarE(vars.int64Var, "int64", 0, "int64 value")
	c.Int8VarE(vars.int8Var, "int8", 0, "int8 value")
	c.QueryStringVarE(vars.queryStringVar, "query-string", map[string]string{}, "query string value")
	c.StringVarE(vars.stringVar, "string", "", "string value")
	c.UintVarE(vars.uintVar, "uint", 0, "uint value")
	c.Uint16VarE(vars.uint16Var, "uint16", 0, "uint16 value")
	c.Uint32VarE(vars.uint32Var, "uint32", 0, "uint32 value")
	c.Uint64VarE(vars.uint64Var, "uint64", 0, "uint64 value")
	c.Uint8VarE(vars.uint8Var, "uint8", 0, "uint8 value")

	testEnvValue(t, c, vars)
}

func TestGlobalEnvValue(t *testing.T) {
	conf.Global = conf.NewConfigurator("name", conf.ContinueOnError)
	vars := new(valueVars)

	vars.boolVar = conf.BoolE("bool", false, "bool value")
	vars.durationVar = conf.DurationE("duration", 0, "time.Duration value")
	vars.float32Var = conf.Float32E("float32", 0, "float32 value")
	vars.float64Var = conf.Float64E("float64", 0, "float64 value")
	vars.intVar = conf.IntE("int", 0, "int value")
	vars.int16Var = conf.Int16E("int16", 0, "int16 value")
	vars.int32Var = conf.Int32E("int32", 0, "int32 value")
	vars.int64Var = conf.Int64E("int64", 0, "int64 value")
	vars.int8Var = conf.Int8E("int8", 0, "int8 value")
	vars.queryStringVar = conf.QueryStringE("query-string", map[string]string{}, "query string value")
	vars.stringVar = conf.StringE("string", "", "string value")
	vars.uintVar = conf.UintE("uint", 0, "uint value")
	vars.uint16Var = conf.Uint16E("uint16", 0, "uint16 value")
	vars.uint32Var = conf.Uint32E("uint32", 0, "uint32 value")
	vars.uint64Var = conf.Uint64E("uint64", 0, "uint64 value")
	vars.uint8Var = conf.Uint8E("uint8", 0, "uint8 value")

	testEnvValue(t, conf.Global, vars)
}

func TestGlobalEnvValueVar(t *testing.T) {
	conf.Global = conf.NewConfigurator("name", conf.ContinueOnError)
	vars := newValueVars()

	conf.BoolVarE(vars.boolVar, "bool", false, "bool value")
	conf.DurationVarE(vars.durationVar, "duration", 0, "time.Duration value")
	conf.Float32VarE(vars.float32Var, "float32", 0, "float32 value")
	conf.Float64VarE(vars.float64Var, "float64", 0, "float64 value")
	conf.IntVarE(vars.intVar, "int", 0, "int value")
	conf.Int16VarE(vars.int16Var, "int16", 0, "int16 value")
	conf.Int32VarE(vars.int32Var, "int32", 0, "int32 value")
	conf.Int64VarE(vars.int64Var, "int64", 0, "int64 value")
	conf.Int8VarE(vars.int8Var, "int8", 0, "int8 value")
	conf.QueryStringVarE(vars.queryStringVar, "query-string", map[string]string{}, "query string value")
	conf.StringVarE(vars.stringVar, "string", "", "string value")
	conf.UintVarE(vars.uintVar, "uint", 0, "uint value")
	conf.Uint16VarE(vars.uint16Var, "uint16", 0, "uint16 value")
	conf.Uint32VarE(vars.uint32Var, "uint32", 0, "uint32 value")
	conf.Uint64VarE(vars.uint64Var, "uint64", 0, "uint64 value")
	conf.Uint8VarE(vars.uint8Var, "uint8", 0, "uint8 value")

	testEnvValue(t, conf.Global, vars)
}

func testEnvValue(t *testing.T, c *conf.Configurator, vars *valueVars) {
	arguments := []string{}

	environment := map[string]string{
		"BOOL":         "true",
		"DURATION":     "1s",
		"FLOAT32":      "172e12",
		"FLOAT64":      "2718e28",
		"INT":          "22",
		"INT16":        "16",
		"INT32":        "32",
		"INT64":        "64",
		"INT8":         "8",
		"QUERY_STRING": "key=value&key2=value2",
		"STRING":       "string",
		"UINT":         "22",
		"UINT16":       "16",
		"UINT32":       "32",
		"UINT64":       "64",
		"UINT8":        "8",
	}

	err := c.Parse(arguments, environment)

	if err != nil {
		t.Fatal("Parse is expected to return a nil (non-error) value")
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

	if (*vars.queryStringVar)["key"] != "value" || (*vars.queryStringVar)["key2"] != "value2" {
		t.Error("query string var should be `key=value&key2=value2`, got: ", *vars.queryStringVar)
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
