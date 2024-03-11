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
