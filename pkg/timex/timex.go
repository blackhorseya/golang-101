package timex

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// GetRandomDuration returns a random duration between min and max.
func GetRandomDuration(min, max int) time.Duration {
	// Create a buffer to store the random bytes
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// Convert the random bytes to an integer
	randomInt := binary.BigEndian.Uint64(b)

	// Scale the integer to the desired range and convert to a time.Duration
	randomDuration := min + int(randomInt%(uint64(max-min+1)))
	return time.Duration(randomDuration) * time.Second
}
