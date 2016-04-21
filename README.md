# go-env

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com/sakeven/go-env)
[![GoDoc](https://img.shields.io/badge/GoDoc-API%20Documentation-blue.svg?style=flat)](http://godoc.org/github.com/sakeven/go-env)

```
go-env is a go library for using environment variable to configure your project.
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
        Num  int    `env:"INNER_NUMBER"`
        Str  string `env:"INNER_HELLO"`
        Bool bool   `env:"INNER_YES"`
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