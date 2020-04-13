// doc http://godoc.org/github.com/deckarep/golang-set

package xset

import "github.com/deckarep/golang-set"

type (
	Iterator = mapset.Iterator
	Set = mapset.Set
	OrderedPair = mapset.OrderedPair
)

var (
	NewSet                      = mapset.NewSet
	NewSetWith                  = mapset.NewSetWith
	NewSetFromSlice             = mapset.NewSetFromSlice
	NewThreadUnsafeSet          = mapset.NewThreadUnsafeSet
	NewThreadUnsafeSetFromSlice = mapset.NewThreadUnsafeSetFromSlice
)
