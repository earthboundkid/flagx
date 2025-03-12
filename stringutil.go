package flagx

import (
	"iter"
	"strings"
)

func joinFunc[T any](seq iter.Seq[T], sep string, f func(T) string) string {
	var buf strings.Builder
	first := true
	for v := range seq {
		if !first {
			buf.WriteString(sep)
		} else {
			first = false
		}
		s := f(v)
		buf.WriteString(s)
	}
	return buf.String()
}
