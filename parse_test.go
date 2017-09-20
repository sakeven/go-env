package env

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

var envSet = LoadSet()

func Test_Int(t *testing.T) {
	g := Goblin(t)

	g.Describe("Int binding", func() {
		g.It("Should get 2 ", func() {
			os.Unsetenv("TEST")
			envSet.Reload()

			test := envSet.Int("TEST", 2)
			g.Assert(test).Equal(2)
		})

		g.It("Should get 1", func() {
			os.Setenv("TEST", "1")
			envSet.Reload()

			test := envSet.Int("TEST", 2)
			g.Assert(test).Equal(1)
		})
	})
}

func Test_Bool(t *testing.T) {
	g := Goblin(t)

	g.Describe("Bool binding", func() {
		g.It("Should get true ", func() {
			os.Unsetenv("TEST")
			envSet.Reload()

			test := envSet.Bool("TEST", true)
			g.Assert(test).Equal(true)
		})

		g.It("Should get false", func() {
			os.Setenv("TEST", "0")
			envSet.Reload()

			test := envSet.Bool("TEST", true)
			g.Assert(test).Equal(false)
		})
	})
}

func Test_String(t *testing.T) {
	g := Goblin(t)

	g.Describe("String binding", func() {
		g.It("Should get hello ", func() {
			os.Unsetenv("TEST")
			envSet.Reload()

			test := envSet.String("TEST", "hello")
			g.Assert(test).Equal("hello")
		})

		g.It("Should get world", func() {
			os.Setenv("TEST", "world")
			envSet.Reload()

			test := envSet.String("TEST", "hello")
			g.Assert(test).Equal("world")
		})
	})
}
