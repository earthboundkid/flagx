package flagx

import (
	"cmp"
	"errors"
	"flag"
	"fmt"
	"slices"
)

// MustHave is a convenience function that checks that the named flags
// were set on fs. Missing flags are treated with the policy of
// fs.ErrorHandling(): ExitOnError, ContinueOnError, or PanicOnError.
// Returned errors can be unpacked by Missing.
//
// If nil, fs defaults to flag.CommandLine.
func MustHave(fs *flag.FlagSet, names ...string) error {
	var missing missingFlagsError
	for f, seen := range All(fs) {
		if !seen && slices.Contains(names, f.Name) {
			missing = append(missing, f)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return handleErr(fs, missing)
}

// Missing returns a slice of required flags missing from an error returned by MustHave.
func Missing(err error) []*flag.Flag {
	var missing missingFlagsError
	errors.As(err, &missing)
	return missing
}

// missingFlagsError is the error type returned by MustHave.
type missingFlagsError []*flag.Flag

func (missing missingFlagsError) Error() string {
	if len(missing) == 0 {
		return "missingFlagsError<empty>"
	}
	if len(missing) == 1 {
		return fmt.Sprintf("missing required flag: %s", missing[0].Name)
	}
	return fmt.Sprintf("missing %d required flags: %s",
		len(missing),
		joinFunc(slices.Values(missing), ", ", func(f *flag.Flag) string {
			return f.Name
		}))
}

// MustHaveArgs is a convenience function that checks that fs.NArg()
// is within the bounds min and max (inclusive). Use max -1 to indicate
// no maximum value. MustHaveArgs uses the policy of fs.ErrorHandling():
// ExitOnError, ContinueOnError, or PanicOnError.
//
// If nil, fs defaults to flag.CommandLine.
func MustHaveArgs(fs *flag.FlagSet, min, max int) error {
	fs = cmp.Or(fs, flag.CommandLine)
	noMax := max < 0
	if max < min && !noMax {
		panic("mismatched arguments to MustHaveArgs")
	}
	n := fs.NArg()
	var err error
	switch {
	case n >= min && (noMax || n <= max):
		return nil
	case min == max && min != 1:
		err = fmt.Errorf("must have %d args; got %d", min, n)
	case min == max:
		err = fmt.Errorf("must have 1 arg; got %d", n)
	case n < min && noMax:
		err = fmt.Errorf("must have at least %d args; got %d", min, n)
	default:
		err = fmt.Errorf("must have between %d and %d args; got %d", min, max, n)
	}
	return handleErr(fs, err)
}
