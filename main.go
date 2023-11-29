package main

import tinymt32 "github.com/aidenfoxivey/gomersenne/mt"

func main() {
	var rng tinymt32.TinyMT32
	rng.Init(4)

	println(rng.Generate_uint32())
}
