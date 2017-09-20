// These examples demonstrate more intricate uses of the env package.
package main

import (
	"log"
	"os"

	"github.com/sakeven/go-env"
)

func assert(expression bool) {
	if !expression {
		panic("Error expression false")
	}
}

func parseStruct() {
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

	assert(test.Bool == true)
	assert(test.Str == "world")
	assert(test.Num == 1)
	assert(test.Default == 12)
	assert(test.Skip == "")
	assert(test.Inner != nil)
	assert(test.Inner.Num == 3)
	assert(test.Inner.Bool == true)
	assert(test.Inner.Str == "inner world")
}

func loadEnvSet() {
	os.Setenv("NUMBER", "1")
	envSet := env.LoadSet()

	n := envSet.Int("NUMBER", 2)

	assert(n == 1)
}

func main() {
	parseStruct()
	loadEnvSet()
	log.Println("Parse env success")
}
