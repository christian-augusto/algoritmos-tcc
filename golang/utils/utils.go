package utils

import (
	"math/rand"
)

func GenerateRandomInt64(seed int64) int64 {
	min := int64(-10_000_000_00)
	max := int64(10_000_000_00)

	rs := rand.NewSource(seed)
	r := rand.New(rs)

	return r.Int63n(max-min) + min
}

func GenerateRandomInt64Array(size int64) []int64 {
	array := make([]int64, size)

	for i := int64(0); i < size; i++ {
		array[i] = GenerateRandomInt64(i)
	}

	return array
}
