package flagx_test

import (
	"flag"
	"fmt"
	"io"

	"github.com/carlmjohnson/flagx"
)

func ExampleChoice_badFlag() {
	fs := flag.NewFlagSet("ExampleChoice", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	var mode string
	fs.Var(flagx.Choice(&mode, "a", "b"), "mode", "mode to run")

	err := fs.Parse([]string{"-mode", "c"})
	fmt.Println(err)
	// Output:
	// invalid value "c" for flag -mode: "c" not in a, b
}

func ExampleChoice_goodFlag() {
	fs := flag.NewFlagSet("ExampleChoice", flag.ContinueOnError)
	var mode string
	fs.Var(flagx.Choice(&mode, "x", "y"), "mode", "mode to run")

	fs.Parse([]string{"-mode", "x"})
	fmt.Println(mode)
	// Output:
	// x
}

func ExampleChoice_defaultValue() {
	fs := flag.NewFlagSet("ExampleChoice", flag.ContinueOnError)
	mode := "none"
	fs.Var(flagx.Choice(&mode, "yes", "no"), "mode", "mode to run")
	fs.Parse([]string{})
	fmt.Println(mode)
	// Output:
	// none
}
