package main

import (
	"fmt"

	"github.com/aidenfoxivey/gotinymt"
)

func main() {
	rng := gotinymt.NewTinyMT(4)
	rng.Init(4)

	for i := 0; i < 4; i++ {
		fmt.Printf("%x\n", rng.Generate_uint32())
	}
}
