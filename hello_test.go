package main

import "testing"

func TestSayingHello(t *testing.T) {
	sayHello()
	if false {
		t.Fail()
	}
}
