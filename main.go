package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/argon2"
)

func main() {
	fmt.Println("parallelism,memory,rounds,time")
	for _, parallelism := range []uint8{4, 8, 12, 16, 32} {
		for _, mem := range []uint32{8, 16, 32, 64} {
			for _, rounds := range []uint32{8, 12, 16, 32} {
				start := time.Now()
				argon2.IDKey([]byte("reknvjkerjkbvrewv"), []byte("brhjkebjberhjvbhjrbfv"), rounds, mem*1024, parallelism, 32)
				duration := time.Since(start)
				fmt.Printf("%d,%d,%d,%f\n", parallelism, mem, rounds, duration.Seconds())
			}
		}
	}
}
