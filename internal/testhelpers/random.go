package testhelpers

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomString() string {
	return strconv.Itoa(newRand().Intn(1000))
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
