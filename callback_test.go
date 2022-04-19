package flagx_test

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/carlmjohnson/flagx"
)

func ExampleCallback_badFlag() {
	fs := flag.NewFlagSet("ExampleCallback", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	const usage = `...`
	var mode string
	flagx.Callback(fs, "mode", mode, usage, func(s string) error {
		if s != strings.ToLower(s) {
			return fmt.Errorf("mode must be lower case")
		}
		mode = s
		return nil
	})

	err := fs.Parse([]string{"-mode", "X"})
	fmt.Println(mode, err)
	// Output:
	// invalid value "X" for flag -mode: mode must be lower case
}

func ExampleCallback_goodFlag() {
	fs := flag.NewFlagSet("ExampleCallback", flag.PanicOnError)
	const usage = `...`
	var mode string
	flagx.Callback(fs, "mode", mode, usage, func(s string) error {
		if s != strings.ToLower(s) {
			return fmt.Errorf("mode must be lower case")
		}
		mode = s
		return nil
	})

	fs.Parse([]string{"-mode", "x"})
	fmt.Println(mode)
	// Output:
	// x
}

func ExampleCallback_defaultValue() {
	fs := flag.NewFlagSet("ExampleCallback", flag.PanicOnError)

	mode := "none"
	flagx.Callback(fs, "mode", mode, "what mode to use", func(s string) error {
		if s != strings.ToLower(s) {
			return fmt.Errorf("mode must be lower case")
		}
		mode = s
		return nil
	})

	fs.Parse([]string{})
	fmt.Println(mode)
	// Output:
	// none
}
