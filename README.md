# go-env
[![Build Status](https://travis-ci.org/sakeven/go-env.svg?branch=master)](https://travis-ci.org/sakeven/go-env)
[![Go Report Card](https://goreportcard.com/badge/github.com/sakeven/go-env)](https://goreportcard.com/report/github.com/sakeven/go-env) 
[![GoDoc](https://godoc.org/github.com/sakeven/go-env?status.svg)](http://godoc.org/github.com/sakeven/go-env)

go-env is a go library for using environment variables to configure your project.

```
Usage:
    // Field appears in env as key "FEILD" and
    // the field is omitted from the object if its value is empty,
    // as defined
    Field int `env:"FEILD"`
    // Field appears in env as key "FEILD" and
    // the field is omitted 1 if its value is empty.
    Field int `env:"FEILD,1"`
    // Field is ignored by this package.
    Field int `env:"-"`
    // The upper key name will be used if it's a non-empty string consisting of
    // only Unicode letters, digits, dollar signs, percent signs, hyphens, underscores and slashes.
    Field int `env:",1"`
```
# Example

```go
// These examples demonstrate more intricate uses of the env package.
package main

import (
    "log"
    "os"

    "github.com/sakeven/go-env"
)

func Assert(expression bool) {
    if !expression {
        log.Fatal("Error expression false")
    }
}

func ParseStruct() {
    type Inner struct {
        Num  int    `env:"NUMBER"`
        Str  string `env:"HELLO"`
        Bool bool   `env:"YES"`
    }

    type Test struct {
        Num     int    `env:"NUMBER"`
        Str     string `env:"HELLO"`
        Bool    bool   `env:"YES"`
        Default int    `env:"DEFAULT,12"`
        Skip    string `env:"-"`
        Inner   *Inner
    }

    os.Setenv("HELLO", "world")
    os.Setenv("NUMBER", "1")
    os.Setenv("YES", "true")
    os.Setenv("SKIP", "skip")

    os.Setenv("INNER_HELLO", "inner world")
    os.Setenv("INNER_NUMBER", "3")
    os.Setenv("INNER_YES", "true")

    test := new(Test)
    err := env.Decode(test)
    if err != nil {
        log.Fatal(err)
    }

    Assert(test.Bool == true)
    Assert(test.Str == "world")
    Assert(test.Num == 1)
    Assert(test.Default == 12)
    Assert(test.Skip == "")
    Assert(test.Inner != nil)
    Assert(test.Inner.Num == 3)
    Assert(test.Inner.Bool == true)
    Assert(test.Inner.Str == "inner world")
}

func LoadEnvSet() {
    os.Setenv("NUMBER", "1")
    envSet := env.Load()

    n := envSet.Int("NUMBER", 2)

    Assert(n == 1)
}

func main() {
    ParseStruct()
    LoadEnvSet()
    log.Println("Parse env success")
}
```
