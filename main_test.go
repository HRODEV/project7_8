package main

import (
	"testing"
)

func TestSimpleExample(t *testing.T) {
	if 1 == 2 {
		t.Errorf("1 should not be equel to 2")
	}
}
