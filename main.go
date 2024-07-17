package main

import (
	// "fmt"
	tinymt32 "github.com/tuck182/gomersenne/mt"
	// "time"
	"flag"
	"os"
)

func main() {
	var rng tinymt32.TinyMT32

	if len(os.Args) == 1 {
		println("Usage:")
		os.Exit(1)
	}

	wordPtr := flag.String("i", "foo.txt", "Newline separated 32-bit integers.")
	flag.Parse()
	println(wordPtr)
	// var seed int
	// var times int
	// fmt.Printf("Enter a seed: ")
	// fmt.Scanf("%d", &seed)

	// rng.Init(uint32(seed))

	// fmt.Printf("Enter a number of times to run the twister: ")
	// fmt.Scanf("%d", &times)

	// for i:=0; i < times; i++ {
	// 	rng.Generate_uint32()
	// }

	// var collected [624]uint32

	// println("Collecting output.")

	// currentTime := time.Now()
	// for i:=0; i < 624; i++ {
	// 	collected[i] = rng.Generate_uint32()
	// }
	// fmt.Printf("%.2f iter/s\n", 624 / time.Now().Sub(currentTime).Seconds())

	println(rng.State_buffer_size())

	// tinymt32.Find_seed([]uint32{1, 2, 3})
}
