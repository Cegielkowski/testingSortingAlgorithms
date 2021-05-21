package main

import (
	crand "crypto/rand"

	"encoding/binary"
	"log"
	"math/rand"
)

func Random() int {
	var src cryptoSource
	rnd := rand.New(src)
	return rnd.Intn(100000) // a truly random number 0 to 999
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
