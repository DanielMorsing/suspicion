package suspicion

import (
	"math/rand"
)

func generateBuf(length int) []byte {
	b := make([]byte, length)
	for i := 0; i < len(b); i++ {
		b[i] = byte(rand.Intn(255))
	}
	return b
}
