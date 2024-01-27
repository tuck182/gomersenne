package mt

import (
	"math"
	"unsafe"
)

const MIN_LOOP = 8
const PRE_LOOP = 8
const TINYMT32_MEXP = 127
const TINYMT32_SH0 = 1
const TINYMT32_SH1 = 10
const TINYMT32_SH8 = 8
const TINYMT32_MASK = 0x7fffffff
const TINYMT32_MUL = float32(1.0 / 16777216.0)

// A 32 bit implementation of the TinyMT mersenne twister
//
// Details found in RFC8682: https://datatracker.ietf.org/doc/rfc8682/
//
// Original website is https://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/TINYMT/
type TinyMT32 struct {
	status [4]uint32
	mat1   uint32
	mat2   uint32
	tmat   uint32
}

func (rng *TinyMT32) periodCertification() {
	if ((rng.status[0] & TINYMT32_MASK) == 0) &&
		(rng.status[1] == 0) &&
		(rng.status[2] == 0) &&
		(rng.status[3] == 0) {
		rng.status[0] = 'T'
		rng.status[1] = 'I'
		rng.status[2] = 'N'
		rng.status[3] = 'Y'
	}
}

// Initializes the tiny mersenne twister.
func (rng *TinyMT32) Init(seed uint32) {
	var i, top uint32 = 1, MIN_LOOP
	rng.status[0] = seed
	rng.status[1] = rng.mat1
	rng.status[2] = rng.mat2
	rng.status[3] = rng.tmat

	for i < top {
		rng.status[i&3] ^= i + 0x6C078965*(rng.status[(i-1)&3]^(rng.status[(i-1)&3]>>30))
		i++
	}
	rng.periodCertification()
	for i = 0; i < PRE_LOOP; i++ {
		rng.NextState()
	}
}

// Advances the tiny mersenne twister to its next state.
func (rng *TinyMT32) NextState() {
	var x, y, a, b uint32

	y = rng.status[3]
	x = (rng.status[0] & TINYMT32_MASK) ^ rng.status[1] ^ rng.status[2]
	x ^= x << TINYMT32_SH0
	y ^= y>>TINYMT32_SH0 ^ x
	rng.status[0] = rng.status[1]
	rng.status[1] = rng.status[2]
	rng.status[2] = x ^ (y << TINYMT32_SH1)
	rng.status[3] = y
	a = -(y & 1) & rng.mat1
	b = -(y & 1) & rng.mat2
	rng.status[1] ^= a
	rng.status[2] ^= b
}

// outputs uint32 from internal state
// shouldn't be called directly
func (rng *TinyMT32) temper() uint32 {
	var t0, t1 uint32

	t0 = rng.status[3]
	t1 = rng.status[0] ^ (rng.status[2] >> TINYMT32_SH8)
	t0 ^= t1
	if (t1 & 1) != 0 {
		t0 ^= rng.tmat
	}
	return t0
}

// outputs float32 from internal state
// shouldn't be called directly
func (rng *TinyMT32) temperConv() float32 {
	var t0, t1, tmp uint32

	t0 = rng.status[3]
	t1 = rng.status[0] ^ (rng.status[2] >> TINYMT32_SH8)
	t0 ^= t1
	if (t1 & 1) != 0 {
		tmp = ((t0 ^ rng.tmat) >> 9) | 0x3f800000
	} else {
		tmp = (t0 >> 9) | 0x3f800000
	}
	return float32(tmp)
}

func (rng *TinyMT32) temperConvOpen() float32 {
	var t0, t1, tmp uint32

	t0 = rng.status[3]
	t1 = rng.status[0] ^ (rng.status[2] >> TINYMT32_SH8)
	t0 ^= t1
	if (t1 & 1) != 0 {
		tmp = ((t0 ^ rng.tmat) >> 9) | 0x3f800001
	} else {
		tmp = (t0 >> 9) | 0x3f800001
	}
	return float32(tmp)
}

func (rng *TinyMT32) Generate_uint32() uint32 {
	rng.NextState()
	return rng.temper()
}

func (rng *TinyMT32) Generate_float() float32 {
	rng.NextState()

	return rng.temperConv() / float32(math.Pow(10, 9))
}

// Returns the state buffer size in terms of bytes.
func (rng *TinyMT32) State_buffer_size() uintptr {
	return unsafe.Sizeof(rng.status)
}
