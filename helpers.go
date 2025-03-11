package flagx

import (
	"cmp"
	"flag"
	"fmt"
	"os"
)

func handleErr(fs *flag.FlagSet, err error) error {
	fs = cmp.Or(fs, flag.CommandLine)
	fmt.Fprintln(fs.Output(), err)
	if fs.Usage != nil {
		fs.Usage()
	}
	switch fs.ErrorHandling() {
	case flag.PanicOnError:
		panic(err)
	case flag.ExitOnError:
		os.Exit(2)
	}
	return err
}
