package codes

import (
	"math/rand"
)

/*------------ generate number with maximum 16 non-zero digits------------*/
func GenCode() uint64 {
	return uint64(rand.Int63n(1e16))
}
