package flagx_test

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/earthboundkid/assert"
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
	be := assert.FailsNow(t)
	// Don't override
	fs := flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	var buf strings.Builder
	fs.SetOutput(&buf)
	a := fs.Int("a", 0, "")
	err := fs.Parse([]string{"-a", "1"})
	be.NilError(err)
	be.Equal(*a, 1)
	// Does not override existing values
	os.Setenv("TEST_ENV_A", "y")
	err = flagx.ParseEnv(fs, "TEST_ENV")
	be.NilError(err)
	output := buf.String()
	be.Falsey(output)

	// Convert kebabs
	fs = flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	buf.Reset()
	fs.SetOutput(&buf)
	kebab := fs.Int("a-b-c", 0, "")
	os.Setenv("TEST_ENV_A_B_C", "1")
	err = flagx.ParseEnv(fs, "TEST_ENV")
	be.NilError(err)
	be.Equal(*kebab, 1)
	output = buf.String()
	be.Falsey(output)

	// With error
	fs = flag.NewFlagSet("ExampleParseEnv", flag.ContinueOnError)
	buf.Reset()
	fs.SetOutput(&buf)
	b := fs.Int("b", 0, "")
	err = fs.Parse(nil)
	be.NilError(err)
	be.Falsey(*b)
	os.Setenv("TEST_ENV_B", "y")
	err = flagx.ParseEnv(fs, "TEST_ENV")
	be.Truthy(err)
	output = buf.String()
	expected := "invalid value \"y\" for flag -b: parse error\nUsage of ExampleParseEnv:\n  -b int\n    \t\n"
	be.Equal(output, expected)
}
