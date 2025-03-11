package flagx

import (
	"cmp"
	"flag"
	"iter"
)

// All returns a sequence
// yielding all flags in fl in lexicographic order
// and boolean indicating whether the flag has been seen yet.
func All(fl *flag.FlagSet) iter.Seq2[*flag.Flag, bool] {
	fl = cmp.Or(fl, flag.CommandLine)
	return func(yield func(*flag.Flag, bool) bool) {
		seenFlags := make(map[*flag.Flag]struct{})
		fl.Visit(func(f *flag.Flag) {
			seenFlags[f] = struct{}{}
		})
		done := false
		fl.VisitAll(func(f *flag.Flag) {
			if done {
				return
			}
			_, ok := seenFlags[f]
			if !yield(f, ok) {
				done = true
			}
		})
	}
}
