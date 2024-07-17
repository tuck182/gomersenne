package main

import (
	"testing"

	tinymt32 "github.com/tuck182/gomersenne/mt"
)

func TestRNG(t *testing.T) {
	var rng tinymt32.TinyMT32
	rng.Init(4)

	want := uint32(1420865856)
	got := rng.Generate_uint32()
	if want != got {
		t.Errorf("rng.Generate_uint32() = %x, want %x", got, want)
	}
}
