package main

import (
	"encoding/binary"
	"fmt"
	"math/rand/v2"
)

type plainRander struct{}

func (plainRander) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = byte(rand.Uint())
	}
	return len(b), nil
}

// always assume it's 16 bytes
type sixteenBytesRander struct{}

func (sixteenBytesRander) Read(b []byte) (int, error) {
	binary.BigEndian.PutUint64(b[:8], rand.Uint64())
	binary.BigEndian.PutUint64(b[8:], rand.Uint64())
	return len(b), nil
}

func main() {
	a := plainRander{}
	b := make([]byte, 16)
	a.Read(b)
	fmt.Println(b)

	a2 := sixteenBytesRander{}
	b2 := make([]byte, 16)
	a2.Read(b2)
	fmt.Println(b2)
}
