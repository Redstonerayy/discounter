package main

import (
	"fmt"
	"math/rand"
)

func newCode() int64 {
	s := rand.Int63n(1e16)
	fmt.Println(s)
	return s
}
