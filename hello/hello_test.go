package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris"

	if got != want {
		t.Error("Got:", got, "Wanted:", want)
	}
}
