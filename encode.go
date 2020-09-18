package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	key, err := strconv.ParseUint(os.Args[1], 0x10, 8)
	if err != nil {
		log.Fatal(err)
	}

	buffer, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	buffer = bytes.Trim(buffer, " \n\t")
	ciphertext := make([]byte, len(buffer))
	for i := range buffer {
		ciphertext[i] = byte(key) ^ buffer[i]
	}

	encodedtext := make([]byte, hex.EncodedLen(len(ciphertext)))
	hex.Encode(encodedtext, ciphertext)

	fmt.Printf("%s\n", encodedtext)
}
