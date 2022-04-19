package flagx_test

import (
	"flag"
	"log"
	"os"

	"github.com/carlmjohnson/flagx"
)

func ExampleLogger() {
	// Redirect Stderr for test
	{
		stderr := os.Stderr
		os.Stderr = os.Stdout
		defer func() {
			os.Stderr = stderr
		}()
	}
	{
		fs := flag.NewFlagSet("ExampleLogger", flag.PanicOnError)
		l := log.New(nil, "myapp ", 00)
		fs.Var(
			flagx.Logger(l, flagx.LogVerbose),
			"verbose",
			`log output`,
		)
		fs.Parse([]string{"-verbose"})

		l.Print("hello log 1")
	}
	{
		fs := flag.NewFlagSet("ExampleLogger", flag.PanicOnError)
		l := log.New(nil, "myapp ", 00)
		fs.Var(
			flagx.Logger(l, flagx.LogSilent),
			"silent",
			`don't log output`,
		)
		fs.Parse([]string{})

		l.Print("hello log 2")
	}
	{
		fs := flag.NewFlagSet("ExampleLogger", flag.PanicOnError)
		l := log.New(nil, "myapp ", 00)
		fs.Var(
			flagx.Logger(l, flagx.LogVerbose),
			"verbose",
			`log output`,
		)
		fs.Parse([]string{"-verbose=false"})

		l.Print("silenced!")
	}
	{
		fs := flag.NewFlagSet("ExampleLogger", flag.PanicOnError)
		l := log.New(nil, "myapp ", 00)
		fs.Var(
			flagx.Logger(l, flagx.LogSilent),
			"silent",
			`don't log output`,
		)
		fs.Parse([]string{"-silent=1"})

		l.Print("silenced!")
	}
	// Output:
	// myapp hello log 1
	// myapp hello log 2
}
