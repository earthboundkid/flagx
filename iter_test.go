package flagx_test

import (
	"flag"
	"io"
	"maps"
	"testing"

	"github.com/carlmjohnson/be"
	"github.com/earthboundkid/flagx/v2"
)

func TestAll(t *testing.T) {
	fs := flag.NewFlagSet("ExampleMustHave", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.String("a", "", "this value must be set")
	fs.String("b", "", "this value must be set")
	fs.String("c", "", "this value is optional")
	fs.Parse([]string{"-a", "set"})
	for f, ok := range flagx.All(fs) {
		be.Equal(t, "a", f.Name)
		be.Equal(t, true, ok)
		break
	}

	m := maps.Collect(flagx.All(fs))
	be.DeepEqual(t, map[*flag.Flag]bool{
		fs.Lookup("a"): true,
		fs.Lookup("b"): false,
		fs.Lookup("c"): false,
	}, m)
}
