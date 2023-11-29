package GoMersenne

import "testing"

func TestRNG(t *testing.T) {
	var rng TinyMT32
	rng.Init(4)

	want := uint32(1420865856)
	got := rng.Generate_uint32()
	if want != got {
		t.Errorf("rng.Generate_uint32() = %x, want %x", got, want)
	}
}

// write unit tests for the other functions
func TestRNG2(t *testing.T) {
	var rng TinyMT32
	rng.Init(4)

	want := uint32(1420865856)
	got := rng.Generate_uint32()
	if want != got {
		t.Errorf("rng.Generate_uint32() = %x, want %x", got, want)
	}
}

func TestRNG3(t *testing.T) {
	var rng TinyMT32
	rng.Init(4)

	want := uint32(1420865856)
	got := rng.Generate_uint32()
	if want != got {
		t.Errorf("rng.Generate_uint32() = %x, want %x", got, want)
	}
}
