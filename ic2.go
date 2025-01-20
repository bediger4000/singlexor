package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	inputFileName := flag.String("f", "", "input file name")
	readHex := flag.Bool("h", false, "read input as ASCII hex-encoded")
	flag.Parse()

	if *inputFileName == "" {
		*inputFileName = flag.Arg(0)
	}

	buffer, err := ioutil.ReadFile(*inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	buffer = bytes.Trim(buffer, " \n\t")
	fmt.Printf("Found %d bytes of hex-encoded ciphertext\n", len(buffer))

	if *readHex {
		ciphertext := make([]byte, hex.DecodedLen(len(buffer)))
		cipherTextLen, err := hex.Decode(ciphertext, buffer)
		if err != nil {
			log.Fatal(err)
		}
		buffer = ciphertext[:cipherTextLen]
	}

	N, ic := indexOfCoincidence(buffer)

	fmt.Printf("%d\t%.05f\n", N, ic)
}

func indexOfCoincidence(buffer []byte) (int, float64) {
	N := 0
	var frequencies [256]int

	for _, r := range buffer {
		frequencies[r]++
		N++
	}

	sum := 0
	for _, freq := range frequencies {
		sum += freq * (freq - 1)
	}

	ic := float64(sum) / float64(N*(N-1))

	return N, ic
}
