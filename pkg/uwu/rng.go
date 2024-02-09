// Contains the internal RNG used for testing.

package uwu

import (
	"math/rand"
	"time"
)

// rng is the internal random number generator.
//
// The purpose of this is for testing; this should get overridden
// in automated test environments.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// fakeRNG is a fakes the RNG for testing purposes. It sets the
// seed to 0 (and as a result, all tests should assume the seed
// is 0).
func fakeRNG() {
	rng = rand.New(rand.NewSource(0))
}
