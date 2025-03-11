package flagx

import (
	"cmp"
	"flag"
	"iter"
)

// All returns a sequence
// yielding all flags in fs in lexicographic order
// and boolean indicating whether the flag has been seen yet.
//
// If nil, fs defaults to flag.CommandLine.
func All(fs *flag.FlagSet) iter.Seq2[*flag.Flag, bool] {
	fs = cmp.Or(fs, flag.CommandLine)
	return func(yield func(*flag.Flag, bool) bool) {
		seenFlags := make(map[*flag.Flag]struct{})
		fs.Visit(func(f *flag.Flag) {
			seenFlags[f] = struct{}{}
		})
		done := false
		fs.VisitAll(func(f *flag.Flag) {
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
