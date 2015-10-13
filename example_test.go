// These examples demonstrate more intricate uses of the env package.
package env_test

import (
	"log"
	"os"

	"github.com/sakeven/go-env"
)

func ExampleDecode() {

	type Inner struct {
		Num  int    `env:"NUMBER"`
		Str  string `env:"HELLO"`
		Bool bool   `env:"YES"`
	}

	type Test struct {
		Num   int    `env:"NUMBER"`
		Str   string `env:"HELLO"`
		Bool  bool   `env:"YES"`
		Inner *Inner
	}

	os.Setenv("HELLO", "world")
	os.Setenv("NUMBER", "1")
	os.Setenv("YES", "true")

	test := new(Test)
	err := env.Decode(test)
	if err != nil {
		log.Fatal(err)
	}

	envSet := env.Load()

	n := envSet.Int("NUMBER", 2)

	log.Println(n)
}
