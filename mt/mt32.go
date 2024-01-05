package mt

type MT32 struct {
	status [624]uint32
	index  uint32
}

func (rng *MT32) initialize_generator(seed uint32) {
	var i, top uint32 = 1, 623
	rng.status[0] = seed

	for i < top {
		rng.status[i] = 0xFFFF & (0x6C078965*(rng.status[i-1]^(rng.status[i-1]>>30)) + i)
		i++
	}
}

func (rng *MT32) extract_number() uint32 {
	if rng.index == 0 {
		rng.generate_numbers()
	}

	y := rng.status[rng.index]

	y = y ^ (y >> 11)
	y = y ^ ((y << 7) & 0x9D2C5680)
	y = y ^ ((y << 15) & 0xEFC60000)
	y = y ^ (y >> 18)

	rng.index = (rng.index + 1) % 624
	return y
}

func (rng *MT32) generate_numbers() {
	var i, y uint32

	for i = 0; i < 624; i++ {
		y = (rng.status[i] & 0x80000000) + (rng.status[(i+1)%624] & 0x7FFFFFFF)
		rng.status[i] = rng.status[(i+397)%624] ^ (y >> 1)
		if (y % 2) != 0 {
			rng.status[i] = rng.status[i] ^ 0x9908B0DF
		}
	}
}

func Find_seed(prior []uint32) {
	println(len(prior))
}
