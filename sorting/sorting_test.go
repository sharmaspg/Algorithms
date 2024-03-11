package sorting_test

import (
	"testing"

	"github.com/sharmaspg/Algorithms/sorting"
)

func TestInsertion(t *testing.T) {
	input := []int64{30, 20, 10}
	wanted := []int64{10, 20, 30}

	result := sorting.Insertion(input)

	if len(result) != len(input) {
		t.Errorf("Wanted %d got %d", len(wanted), len(result))
	}
	for index, val := range result {
		if wanted[index] != val {
			t.Errorf("Wanted %d got %d", wanted[index], val)
		}
	}

}
