package flagx

import (
	"cmp"
	"flag"
	"fmt"
	"os"
)

func handleErr(fl *flag.FlagSet, err error) error {
	fl = cmp.Or(fl, flag.CommandLine)
	fmt.Fprintln(fl.Output(), err)
	if fl.Usage != nil {
		fl.Usage()
	}
	switch fl.ErrorHandling() {
	case flag.PanicOnError:
		panic(err)
	case flag.ExitOnError:
		os.Exit(2)
	}
	return err
}
