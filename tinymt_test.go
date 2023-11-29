package GoTinyMT

import "testing"

func TestHello(t *testing.T) {
	var rng TinyMT32
	rng.Init(4)

	want := uint32(1420865856)
	got := rng.Generate_uint32()
	if want != got {
		t.Errorf("rng.Generate_uint32() = %x, want %x", got, want)
	}
}
