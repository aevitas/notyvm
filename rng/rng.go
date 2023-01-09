package rng

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"log"
	"math/rand"
)

func RandomNumber() int {
	return rand.New(getRand()).Int()
}

func getRand() rand.Source {
	var entropy [8]byte
	_, err := crypto_rand.Read(entropy[:])

	if err != nil {
		log.Fatal("could not read 8 bytes of entropy from crypto rand")
	}

	return rand.NewSource(int64(binary.LittleEndian.Uint64(entropy[:])))
}
