package sorting

func Insertion(A []int64) []int64 {
	for j := 1; j < len(A); j++ {
		i := j - 1
		key := A[j]
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
	return A
}

func MergeSort(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func Selection(A []int64) []int64 {
	var min_index int
	var temp int64
	for i := 0; i < len(A)-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[min_index] {
				min_index = j
			}
		}
		temp = A[i]
		A[i] = A[min_index]
		A[min_index] = temp
	}
	return A
}

func merge(left, right []int64) []int64 {
	result := make([]int64, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
