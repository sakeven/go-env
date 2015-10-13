package env

import (
	"os"
	"testing"

	. "github.com/franela/goblin"
)

func Test_Decode(t *testing.T) {
	g := Goblin(t)

	type Test struct {
		Num  int    `env:"NUMBER"`
		Str  string `env:"HELLO"`
		Bool bool   `env:"YES"`
	}

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

		test := new(Test)
		err := Decode(test)
		if err != nil {
			g.Fail(err)
		}

		g.It("Should get int 1", func() {
			g.Assert(test.Num).Equal(1)
		})

		g.It("Should get inner int 1", func() {
			g.Assert(test.Inner.Num).Equal(1)
		})

		g.It("Should get string world", func() {
			g.Assert(test.Str).Equal("world")
		})

		g.It("Should get inner string world", func() {
			g.Assert(test.Inner.Str).Equal("world")
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

	g.Describe("Decode inner pointer struct", func() {
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

		g.It("Should get inner int 1", func() {
			g.Assert(test.Inner.Num).Equal(1)
		})

		g.It("Should get string world", func() {
			g.Assert(test.Str).Equal("world")
		})

		g.It("Should get inner string world", func() {
			g.Assert(test.Inner.Str).Equal("world")
		})

		g.It("Should get bool true", func() {
			g.Assert(test.Bool).Equal(true)
		})

		g.It("Should get inner bool true", func() {
			g.Assert(test.Inner.Bool).Equal(true)
		})
	})
}
