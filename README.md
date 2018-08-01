# Conf - CLI configuration management

[![Build Status](https://travis-ci.com/goph/conf.svg?branch=master)](https://travis-ci.com/goph/conf)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/conf?style=flat-square)](https://goreportcard.com/report/github.com/goph/conf)
[![GolangCI](https://golangci.com/badges/github.com/goph/conf.svg)](https://golangci.com)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/conf)

**CLI configuration via flags and environment variables.**


## Usage

The interface follows closely the [flag](https://golang.org/pkg/flag)/[pflag](https://github.com/spf13/pflag) packages.

```go
package main

import "github.com/goph/conf"

func main() {
	// Define environment variables and flags using conf.String(), Bool(), Int(), etc.
	var intValue *int = conf.Int("int", 1234, "help message for int")
	
	// If you like, you can bind the environment variable and the flag to a variable using the Var() functions.
	var intVar int
	conf.IntVar(&intVar, "int", 1234, "help message for int")
	
	// Or you can create custom variables that satisfy the Value interface (with pointer receivers)
	// and couple them to variable parsing.
	//
	// For such environment variables/flags, the default value is just the initial value of the variable.
	var intVal conf.Value
	conf.Var(&intVal, "int", "help message for int")
	
	// If you prefer only environment variables or flags as a source for certain values,
	// you can use the above methods with the "E" and "F" suffix.
	intValue = conf.IntF("int", 1234, "help message for int")
	intValue = conf.IntE("int", 1234, "help message for int")
	
	// After all variables are defined.
	conf.Parse()
}
```


## Development

[dep](https://golang.github.io/dep/) is required to install dependencies:

```bash
$ make dep # dep ensure
```

When all coding and testing is done, please run the test suite:

```bash
$ make test # go test
```

For linting we use [GolangCI](https://golangci.com/).

```bash
$ make lint # golangci-lint run
```

You can run the whole suite with:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

This project is heavily inspired by [flag](https://golang.org/pkg/flag)/[pflag](https://github.com/spf13/pflag) packages.
