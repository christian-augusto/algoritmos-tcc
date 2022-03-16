package utils

import (
	"math/rand"
)

func GenerateRandomInt64Array(size int64) []int64 {
	array := make([]int64, size)

	for i := int64(0); i < size; i++ {
		array[i] = GenerateRandomInt64(i)
	}

	return array
}

func GenerateRandomInt64(seed int64) int64 {
	min := int64(0)
	max := int64(20_000_000_00)
	// min := int64(-10)
	// max := int64(10)

	rs := rand.NewSource(seed)
	r := rand.New(rs)

	return r.Int63n(max-min) + min
}
