package codes

import (
	"fmt"
	"math/rand"
)

func GenCode() uint64 {
	a := rand.Int63n(1e16)
	fmt.Println(a)
	s := uint64(a)
	fmt.Println(s)
	return s
}
