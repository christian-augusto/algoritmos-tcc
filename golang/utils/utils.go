package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomIntArray(size int, min int, max int) []int {
	array := make([]int, size)

	for i := int(0); i < size; i++ {
		array[i] = generateRandomInt(int64(i)+time.Now().Unix(), min, max)
	}

	return array
}

func generateRandomInt(seed int64, min int, max int) int {
	rs := rand.NewSource(seed)
	r := rand.New(rs)

	return r.Intn(max-min) + min
}

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}

	return b
}
