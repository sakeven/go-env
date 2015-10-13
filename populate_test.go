package env

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

type Test struct {
	Num  int    `env:"NUMBER"`
	Str  string `env:"HELLO"`
	Bool bool   `env:"YES"`
}

func Test_Decode(t *testing.T) {
	g := Goblin(t)

	g.Describe("Decode test", func() {
		os.Setenv("HELLO", "world")
		os.Setenv("NUMBER", "1")
		os.Setenv("YES", "true")

		test := new(Test)
		err := Decode(test)
		if err != nil {
			g.Fail(err)
		}

		g.It("Should get int 1", func() {
			g.Assert(test.Num).Equal(1)
		})

		g.It("Should get string world", func() {
			g.Assert(test.Str).Equal("world")
		})

		g.It("Should get bool true", func() {
			g.Assert(test.Bool).Equal(true)
		})
	})
}
