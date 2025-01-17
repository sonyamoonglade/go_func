package go_func

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.SliceIsSorted.
func AreSorted[T constraints.Ordered](ss []T) bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}
