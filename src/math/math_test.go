package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 4)

	if result != 5 {
		t.Errorf("Add(1, 4) FAILED. Expected %d, got %d \n", 5, result)
	} else {
		t.Logf("Add(1, 4) PASSED, Expected %d, got %d", 5, result)
	}
}
