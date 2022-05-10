package flagx_test

import (
	"flag"
	"fmt"

	"github.com/carlmjohnson/flagx"
)

func ExampleBoolFunc() {
	fs := flag.NewFlagSet("ExampleParseEnv", flag.PanicOnError)
	flagx.BoolFunc(fs, "call-me", "", func() error {
		fmt.Println("called!")
		return nil
	})
	flagx.BoolFunc(fs, "dont-call-me", "", func() error {
		fmt.Println("not called!")
		return nil
	})
	fs.Parse([]string{"-call-me", "-dont-call-me=false"})
	// Output:
	// called!
}
