package b

import "testing"

func TestBanana(t *testing.T) {
	if Banana != 2 {
		t.Fail()
	}
}
