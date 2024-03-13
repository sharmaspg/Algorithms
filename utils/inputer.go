package utils

func GetArray(size int64) []int64 {
	result := []int64{}
	x := int64(4555486587645876987)
	for i := int64(0); i < size; i += 1 {
		result = append(result, x)
		x = x - 1

	}
	return result
}

func GetInputForGraph() [][]int64 {
	result := [][]int64{}
	for i := 5000; i < 55000; i += 100 {
		result = append(result, GetArray(int64(i)))
	}

	return result
}
