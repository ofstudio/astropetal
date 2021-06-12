package timing

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Approx(base, delta time.Duration) time.Duration {
	return base + time.Duration(rand.Int63n(int64(delta)))*2 - delta
}
