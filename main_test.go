package main

import "testing"

func TestHello(t *testing.T) {
    if hello() != "Hello" {
        t.Error("Testing error")
    }
}