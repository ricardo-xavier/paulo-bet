package utils

import (
    "testing"
)

func TestCript(t *testing.T) {
    token := BuildToken("ricardo.xavier")
    got := CheckToken("ricardo.xavier", token)
    want := true
    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}
