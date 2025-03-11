package flagx_test

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/earthboundkid/flagx/v2"
)

func ExampleMustHave_missingFlag() {
	fs := flag.NewFlagSet("ExampleMustHave", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	fs.String("a", "", "this value must be set")
	fs.String("b", "", "this value must be set")
	fs.String("c", "", "this value is optional")
	fs.Parse([]string{"-a", "set"})
	err := flagx.MustHave(fs, "a", "b")
	fmt.Println("Missing:", flagx.Missing(err))
	// Output:
	// missing required flag: b
	// Usage of ExampleMustHave:
	//   -a string
	//     	this value must be set
	//   -b string
	//     	this value must be set
	//   -c string
	//     	this value is optional
	// Missing: [b]
}

func ExampleMustHave_noMissingFlag() {
	fs := flag.NewFlagSet("ExampleMustHave", flag.PanicOnError)
	fs.String("a", "", "this value must be set")
	fs.String("b", "", "this value must be set")
	fs.String("c", "", "this value is optional")
	fs.Parse([]string{"-a", "set", "-b", "set"})
	flagx.MustHave(fs, "a", "b")
	// Output:
}

func ExampleMustHaveArgs_wrongNumber() {
	var buf strings.Builder
	defer func() {
		recover()
		fmt.Println(buf.String())
	}()

	fs := flag.NewFlagSet("ExampleMustHaveArgs", flag.PanicOnError)
	fs.SetOutput(&buf)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage:\n\tExampleMustHaveArgs [optional arg]")
	}

	fs.Parse([]string{"--", "one", "two"})
	flagx.MustHaveArgs(fs, 0, 1)
	// Output:
	// must have between 0 and 1 args; got 2
	// Usage:
	// 	ExampleMustHaveArgs [optional arg]
}

func ExampleMustHaveArgs_correctNumber() {
	fs := flag.NewFlagSet("ExampleMustHave", flag.PanicOnError)
	fs.String("a", "", "an option")
	fs.Parse([]string{"--", "-a", "-b", "-c"})
	flagx.MustHaveArgs(fs, 3, 3)
	// Output:
}
