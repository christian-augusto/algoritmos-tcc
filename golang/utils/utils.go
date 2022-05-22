package utils

import (
	"math/rand"
)

func GenerateRandomIntArray(size int) []int {
	array := make([]int, size)

	for i := int(0); i < size; i++ {
		array[i] = GenerateRandomInt(i)
	}

	return array
}

func GenerateRandomInt(seed int) int {
	min := int(0)
	max := int(20_000_000_00)

	rs := rand.NewSource(int64(seed))
	r := rand.New(rs)

	return r.Intn(max-min) + min
}
