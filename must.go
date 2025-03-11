package flagx

import (
	"cmp"
	"errors"
	"flag"
	"fmt"
	"slices"
	"strings"
)

// MustHave is a convenience function that checks that the named flags
// were set on fl. Missing flags are treated with the policy of
// fl.ErrorHandling(): ExitOnError, ContinueOnError, or PanicOnError.
// Returned errors can be unpacked by Missing.
//
// If nil, fl defaults to flag.CommandLine.
func MustHave(fl *flag.FlagSet, names ...string) error {
	var missing missingFlagsError
	for f, seen := range All(fl) {
		if !seen && slices.Contains(names, f.Name) {
			missing = append(missing, f.Name)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return handleErr(fl, missing)
}

// Missing returns a slice of required flags missing from an error returned by MustHave.
func Missing(err error) []string {
	var missing missingFlagsError
	errors.As(err, &missing)
	return missing
}

// missingFlagsError is the error type returned by MustHave.
type missingFlagsError []string

func (missing missingFlagsError) Error() string {
	if len(missing) == 0 {
		return "missingFlagsError<empty>"
	}
	if len(missing) == 1 {
		return fmt.Sprintf("missing required flag: %s", missing[0])
	}
	return fmt.Sprintf("missing %d required flags: %s",
		len(missing), strings.Join(missing, ", "))
}

// MustHaveArgs is a convenience function that checks that fl.NArg()
// is within the bounds min and max (inclusive). Use max -1 to indicate
// no maximum value. MustHaveArgs uses the policy of  fl.ErrorHandling():
// ExitOnError, ContinueOnError, or PanicOnError.
//
// If nil, fl defaults to flag.CommandLine.
func MustHaveArgs(fl *flag.FlagSet, min, max int) error {
	fl = cmp.Or(fl, flag.CommandLine)
	noMax := max < 0
	if max < min && !noMax {
		panic("mismatched arguments to MustHaveArgs")
	}
	n := fl.NArg()
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
	return handleErr(fl, err)
}
