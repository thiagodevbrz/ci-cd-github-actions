package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 3)

	if result != 4 {
		t.Errorf("Add(1, 3) FAILED. Expected %d, got %d \n", 4, result)
	} else {
		t.Logf("Add(1, 3) PASSED, Expected %d,, got %d", 4, result)
	}
}
