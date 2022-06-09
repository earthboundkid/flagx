package flagx

import (
	"errors"
	"flag"
	"strconv"
)

// BoolFunc defines a flag with the specified name and usage string.
// Each time the flag is set with a truthy value, fn is called.
// If fn returns a non-nil error, it will be treated as a flag value parsing error.
func BoolFunc(fs *flag.FlagSet, name, usage string, fn func() error) {
	fs = flagOrDefault(fs)
	fs.Var(boolFunc(fn), name, usage)
}

type boolFunc func() error

func (f boolFunc) IsBoolFlag() bool {
	return true
}

func (f boolFunc) String() string {
	return ""
}

func (f boolFunc) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return errors.New("parse error")
	}
	if b {
		return f()
	}
	return nil
}
