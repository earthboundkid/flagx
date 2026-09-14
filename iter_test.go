package flagx_test

import (
	"flag"
	"io"
	"maps"
	"testing"

	"github.com/earthboundkid/assert"
	"github.com/earthboundkid/flagx/v2"
)

func TestAll(t *testing.T) {
	tt := assert.Continues(t)
	fs := flag.NewFlagSet("ExampleMustHave", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.String("a", "", "this value must be set")
	fs.String("b", "", "this value must be set")
	fs.String("c", "", "this value is optional")
	fs.Parse([]string{"-a", "set"})
	for f, ok := range flagx.All(fs) {
		tt.Equal(f.Name, "a").True(ok)
		break
	}

	tt.True(maps.Equal(
		maps.Collect(flagx.All(fs)),
		map[*flag.Flag]bool{
			fs.Lookup("a"): true,
			fs.Lookup("b"): false,
			fs.Lookup("c"): false,
		}))
}
