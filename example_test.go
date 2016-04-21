// These examples demonstrate more intricate uses of the env package.
package env_test

import (
	"log"
	"os"
	"testing"

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

	Assert(n == 2)
}

func Examples(t *testing.T) {
	ParseStruct()
	LoadEnvSet()
}
