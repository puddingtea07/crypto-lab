package main

import (
	"crypto-lab/vigenere"
	"fmt"
	"os"
)

func vigenerEnc() {
	key := os.Args[1]
	filename := os.Args[2]
	fmt.Println(vigenere.Encrypt(key, filename))
}

func main() {
	vigenerEnc()
}