package utils

import (
	"math/rand"
)

func GetArray(size int64) []int64 {
	result := []int64{}

	for i := int64(0); i < size; i += 1 {
		result = append(result, rand.Int63n(int64(4555486587645876)))

	}
	return result
}

func GetInputForGraph() [][]int64 {
	result := [][]int64{}
	for i := 1000; i < 11000; i += 100 {
		result = append(result, GetArray(int64(i)))
	}

	return result
}
