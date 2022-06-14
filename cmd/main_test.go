package main

import "testing"

func Test_Something(t *testing.T) {
	got := Add(4, 6)
    want := 9

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}