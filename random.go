package urandom

import (
	"math/rand"
	"sync"
	"time"
)

type random struct {
	stateLock *sync.Mutex
	rand      *rand.Rand
}

var seed *random

func init() {
	seed = &random{
		rand: rand.New(
			rand.NewSource(time.Now().UnixNano()),
		),
		stateLock: &sync.Mutex{},
	}
}

func (r *random) Intn(n int) int {
	r.stateLock.Lock()
	defer r.stateLock.Unlock()
	return r.rand.Intn(n)
}

func (r *random) Float64() float64 {
	r.stateLock.Lock()
	defer r.stateLock.Unlock()
	return r.rand.Float64()
}
