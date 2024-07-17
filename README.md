# Tiny Mersenne Twister Implementation in Go

This project is a minimalistic implementation of the Mersenne Twister
pseudorandom number generator (PRNG) in Go. The Mersenne Twister algorithm,
developed by Makoto Matsumoto and Takuji Nishimura, is a widely-used PRNG known
for its long period and excellent statistical properties.

[Original TinyMT Page](http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/TINYMT/)

## Features

- Efficient and compact implementation of the Tiny Mersenne Twister algorithm in Go.
- Provides high-quality pseudorandom numbers with a period of 2^127-1.
- Supports both seeding with a single value and seeding with an array of values.
- Easy integration with existing Go projects.
- Standard library Go.

## How does prediction work?

PRNGs (Pseudo Random Number Generators) come in different sorts. Similar to hash
functions, some have the use of being "utility" focused, whereas others are
cryptographically secure.

In particular, f

## Usage

1. Start by installing Go and setting up your development environment.

2. Add the `mersennetwister` package to your project:

   ```bash
   go get github.com/tuck182/gomersenne
   ```

3. Import the package in your code:

   ```go
   import (
      "fmt"

      "github.com/tuck182/gomersenne"
   )
   ```

4. Seed the Mersenne Twister generator with an initial value or an array of values:

   ```go
   seed := uint32(12345)
   mt := TinyMT32.Init(seed)

   // Or with an array of values
   seedArray := []uint32{123, 456, 789}
   mt := TinyMT32.InitArray(seedArray)
   ```

5. Generate pseudorandom numbers:

   ```go
   randNum := mt.Generate_uint32() // Generate a 32-bit unsigned random number
   ```

## License

This project is licensed under the 3 clause BSD license.
