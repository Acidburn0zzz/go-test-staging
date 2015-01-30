package a

import "testing"

func TestApple(t *testing.T) {
	if Apple != 1 {
		t.Fail()
	}
}
