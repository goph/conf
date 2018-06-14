package conf_test

import "time"

type valueVars struct {
	boolVar        *bool
	durationVar    *time.Duration
	float32Var     *float32
	float64Var     *float64
	intVar         *int
	int16Var       *int16
	int32Var       *int32
	int64Var       *int64
	int8Var        *int8
	queryStringVar *map[string]string
	stringVar      *string
	uintVar        *uint
	uint16Var      *uint16
	uint32Var      *uint32
	uint64Var      *uint64
	uint8Var       *uint8
}

func newValueVars() *valueVars {
	return &valueVars{
		boolVar:        new(bool),
		durationVar:    new(time.Duration),
		float32Var:     new(float32),
		float64Var:     new(float64),
		intVar:         new(int),
		int16Var:       new(int16),
		int32Var:       new(int32),
		int64Var:       new(int64),
		int8Var:        new(int8),
		queryStringVar: new(map[string]string),
		stringVar:      new(string),
		uintVar:        new(uint),
		uint16Var:      new(uint16),
		uint32Var:      new(uint32),
		uint64Var:      new(uint64),
		uint8Var:       new(uint8),
	}
}
