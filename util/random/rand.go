package random

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func RandIntInterval(min int, max int, nums int) []int {
	result := []int{}

	for i := 0; i < nums; i++ {
		result = append(result, rand.Intn(max-min+1)+min)
	}

	return result
}
