package c

import "testing"

func TestCarrot(t *testing.T) {
	if Carrot != 3 {
		t.Fail()
	}
}
