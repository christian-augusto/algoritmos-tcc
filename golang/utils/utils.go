package utils

import (
	"algoritmos-tcc-golang/constants"
	"math/rand"
	"time"
)

func GenerateRandomIntArray(size int) []int {
	array := make([]int, size)

	for i := int(0); i < size; i++ {
		array[i] = GenerateRandomInt(int64(i) + time.Now().Unix())
	}

	return array
}

func GenerateRandomInt(seed int64) int {
	min := constants.RANDOM_INT_MIN
	max := constants.RANDOM_INT_MAX

	rs := rand.NewSource(seed)
	r := rand.New(rs)

	return r.Intn(max-min) + min
}
