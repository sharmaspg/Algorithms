package utils_test

import (
	"testing"

	"github.com/sharmaspg/Algorithms/utils"
)

func TestGetArray(t *testing.T) {
	want := 30
	result := utils.GetArray(30)
	if want != len(result) {
		t.Errorf("Wanted %d got %d", want, len(result))
	}
}
