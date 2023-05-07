package helper

import (
	"math/rand"
	"time"
)

type Sos struct {
	Element1 string
	Element2 string
	Element3 string
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
