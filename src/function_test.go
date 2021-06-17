package src

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(1,2,3)
	if sum != 6 {
		t.Errorf("src Sum error")
	}
}
