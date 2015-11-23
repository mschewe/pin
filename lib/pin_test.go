package pin

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func generateRandomDigit() uint64 {
	i, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	return i.Uint64()
}

func Generate(length int) string {
	var pin string

	for i := 0; i < length; i++ {
		rand := generateRandomDigit()
		digit := strconv.FormatUint(rand, 10)
		pin += string(digit)
	}

	return pin
}
