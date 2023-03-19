package random

import (
	"math/rand"
	"time"
)

func Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
