package flagx_test

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/earthboundkid/flagx/v2"
)

func ExampleParseEnv() {
	fs := flag.NewFlagSet("ExampleParseEnv", flag.PanicOnError)
	a := fs.Int("a", 0, "")
	b := fs.Int("b", 0, "")
	fs.Parse([]string{"-a", "1"})

	os.Setenv("TEST_ENV_A", "2")
	os.Setenv("TEST_ENV_B", "3")
	flagx.ParseEnv(fs, "test-env")

	// Does not override existing values
	fmt.Println("a", *a)
	// Does get new values from env
	fmt.Println("b", *b)
	// Output:
	// a 1
	// b 3
}

func TestParseEnv(t *testing.T) {
	// Don't override
	fs := flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	var buf strings.Builder
	fs.SetOutput(&buf)
	a := fs.Int("a", 0, "")
	err := fs.Parse([]string{"-a", "1"})
	be.NilErr(t, err)
	be.Equal(t, 1, *a)
	// Does not override existing values
	os.Setenv("TEST_ENV_A", "y")
	err = flagx.ParseEnv(fs, "TEST_ENV")
	be.NilErr(t, err)
	output := buf.String()
	be.Zero(t, output)

	// Convert kebabs
	fs = flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	buf.Reset()
	fs.SetOutput(&buf)
	kebab := fs.Int("a-b-c", 0, "")
	os.Setenv("TEST_ENV_A_B_C", "1")
	err = flagx.ParseEnv(fs, "TEST_ENV")
	be.NilErr(t, err)
	be.Equal(t, 1, *kebab)
	output = buf.String()
	be.Zero(t, output)

	// With error
	fs = flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	buf.Reset()
	fs.SetOutput(&buf)
	b := fs.Int("b", 0, "")
	err = fs.Parse(nil)
	be.NilErr(t, err)
	be.Zero(t, *b)
	os.Setenv("TEST_ENV_B", "y")
	err = flagx.ParseEnv(fs, "TEST_ENV")
	be.Nonzero(t, err)
	output = buf.String()
	expected := "invalid value \"y\" for flag -b: parse error\nUsage of ExampleParseEnv:\n  -b int\n    \t\n"
	be.Equal(t, expected, output)
}
