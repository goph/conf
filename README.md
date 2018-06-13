# Conf - CLI configuration management

[![Build Status](https://travis-ci.com/goph/conf.svg?branch=master)](https://travis-ci.com/goph/conf)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/conf?style=flat-square)](https://goreportcard.com/report/github.com/goph/conf)
[![GolangCI](https://golangci.com/badges/github.com/goph/conf.svg)](https://golangci.com)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/conf)

**CLI configuration via flags and environment variables.**


## Development

[dep](https://golang.github.io/dep/) is required to install dependencies:

```bash
$ dep ensure
```

When all coding and testing is done, please run the test suite:

``` bash
$ go test
```

For linting we use [GolangCI](https://golangci.com/).
You can run the linter locally using it's [binary version](https://github.com/golangci/golangci-lint#ci-installation):

```bash
$ golangci-lint run
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

This project is heavily inspired by [flag](https://golang.org/pkg/flag)/[pflag](https://github.com/spf13/pflag) packages.
