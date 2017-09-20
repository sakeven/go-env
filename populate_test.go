package env

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func Test_Decode(t *testing.T) {
	g := Goblin(t)

	type Test struct {
		Num      *int   `env:"NUMBER"`
		Str      string `env:"HELLO"`
		Bool     bool   `env:"YES"`
		Default  int    `env:"DEFAULT,12"`
		Skip     string `env:"-"`
		Aphal    string
		unexport string
	}

	g.Describe("Decode test", func() {
		os.Setenv("HELLO", "world")
		os.Setenv("NUMBER", "1")
		os.Setenv("YES", "true")
		os.Setenv("APHAL", "beta")
		os.Setenv("SKIP", "SKIP")

		test := new(Test)
		err := Decode(test)
		if err != nil {
			g.Fail(err)
		}

		g.It("Should get int 1", func() {
			g.Assert(*test.Num).Equal(1)
		})

		g.It("Should get default 12", func() {
			g.Assert(test.Default).Equal(12)
		})

		g.It("Should get string world", func() {
			g.Assert(test.Str).Equal("world")
		})

		g.It("Should get bool true", func() {
			g.Assert(test.Bool).Equal(true)
		})

		g.It("Should get string beta", func() {
			g.Assert(test.Aphal).Equal("beta")
		})
		g.It("Should get empty string", func() {
			g.Assert(test.Skip).Equal("")
		})
	})
}

func Test_Decode_Inner(t *testing.T) {
	g := Goblin(t)

	type Inner struct {
		Num  int    `env:"NUMBER"`
		Str  string `env:"HELLO"`
		Bool bool   `env:"YES"`
	}

	type Test struct {
		Num   int    `env:"NUMBER"`
		Str   string `env:"HELLO"`
		Bool  bool   `env:"YES"`
		Inner Inner
	}

	g.Describe("Decode inner struct", func() {
		os.Setenv("HELLO", "world")
		os.Setenv("NUMBER", "1")
		os.Setenv("YES", "true")

		os.Setenv("INNER_HELLO", "inner world")
		os.Setenv("INNER_NUMBER", "3")
		os.Setenv("INNER_YES", "true")

		test := new(Test)
		err := Decode(test)
		if err != nil {
			g.Fail(err)
		}

		g.It("Should get int 1", func() {
			g.Assert(test.Num).Equal(1)
		})

		g.It("Should get inner int 3", func() {
			g.Assert(test.Inner.Num).Equal(3)
		})

		g.It("Should get string world", func() {
			g.Assert(test.Str).Equal("world")
		})

		g.It("Should get string inner world", func() {
			g.Assert(test.Inner.Str).Equal("inner world")
		})

		g.It("Should get bool true", func() {
			g.Assert(test.Bool).Equal(true)
		})

		g.It("Should get inner bool true", func() {
			g.Assert(test.Inner.Bool).Equal(true)
		})
	})
}

func Test_Decode_Inner_Pointer(t *testing.T) {
	g := Goblin(t)

	type Inner struct {
		Num int `env:"NUMBER"`
	}

	type Test struct {
		Num   int    `env:"NUMBER"`
		Str   string `env:"HELLO"`
		Bool  bool   `env:"YES"`
		Inner *Inner `env:"IN"`
	}

	g.Describe("Decode inner pointer struct", func() {
		os.Setenv("HELLO", "world")
		os.Setenv("NUMBER", "1")
		os.Setenv("YES", "true")

		os.Setenv("IN_NUMBER", "3")

		test := new(Test)
		err := Decode(test)
		if err != nil {
			g.Fail(err)
		}

		g.It("Should get int 1", func() {
			g.Assert(test.Num).Equal(1)
		})

		g.It("Should get inner int 3", func() {
			g.Assert(test.Inner.Num).Equal(3)
		})

		g.It("Should get string world", func() {
			g.Assert(test.Str).Equal("world")
		})

		g.It("Should get bool true", func() {
			g.Assert(test.Bool).Equal(true)
		})
	})
}
