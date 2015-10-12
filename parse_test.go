package env

import (
    "os"
    "testing"

    . "github.com/franela/goblin"
)

func Test_Int(t *testing.T) {
    g := Goblin(t)

    g.Describe("Int binding", func() {
        g.It("Should get 2 ", func() {
            os.Unsetenv("TEST")
            Load()

            test := Int("TEST", 2)
            g.Assert(test).Equal(2)
        })

        g.It("Should get 1", func() {
            os.Setenv("TEST", "1")
            Load()

            test := Int("TEST", 2)
            g.Assert(test).Equal(1)
        })
    })
}

func Test_Bool(t *testing.T) {
    g := Goblin(t)

    g.Describe("Bool binding", func() {
        g.It("Should get true ", func() {
            os.Unsetenv("TEST")
            Load()

            test := Bool("TEST", true)
            g.Assert(test).Equal(true)
        })

        g.It("Should get false", func() {
            os.Setenv("TEST", "0")
            Load()

            test := Bool("TEST", true)
            g.Assert(test).Equal(false)
        })
    })
}

func Test_String(t *testing.T) {
    g := Goblin(t)

    g.Describe("String binding", func() {
        g.It("Should get hello ", func() {
            os.Unsetenv("TEST")
            Load()

            test := String("TEST", "hello")
            g.Assert(test).Equal("hello")
        })

        g.It("Should get world", func() {
            os.Setenv("TEST", "world")
            Load()

            test := String("TEST", "hello")
            g.Assert(test).Equal("world")
        })
    })
}
