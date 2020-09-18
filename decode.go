package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func main() {
	inputFileName := flag.String("f", "", "input file name")
	allowableErrors := flag.Int("e", 0, "allowable non-ASCII bytes")
	maxOutErrors := flag.Bool("i", false, "don't use non-ASCII byte count criteria")
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

	ciphertext := make([]byte, hex.DecodedLen(len(buffer)))
	cipherTextLen, err := hex.Decode(ciphertext, buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d bytes of ciphertext\n", cipherTextLen)

	if *maxOutErrors {
		*allowableErrors = cipherTextLen + 1
	}

	var bestAngle float64
	bestAngle = math.MaxFloat32
	var bestKey byte

	// amusingly, keyByte has to have type int to allow the loop
	// criteria to work. An index of type byte rolls over from 255 to 0,
	// so even "<= 255" won't work
	for keyByte := 0; keyByte < 256; keyByte++ {
		badByteCnt := singleXor(ciphertext, byte(keyByte))
		fmt.Printf("key %02x, errors %d\n", keyByte, badByteCnt)
		if badByteCnt < *allowableErrors {
			printXor(ciphertext, byte(keyByte))
			v := createVector(ciphertext, byte(keyByte))
			angle := vectorAngle(v, englishVector)
			if angle < bestAngle {
				bestAngle = angle
				bestKey = byte(keyByte)
			}
		}
	}
	fmt.Printf("Best key %02x\n", bestKey)
	printXor(ciphertext, bestKey)
}

// singleXor does an XOR of all bytes in buffer with the key byte,
// returning a count of non-ASCII bytes when decoded this way
func singleXor(buffer []byte, key byte) int {
	errorCount := 0
	for i := range buffer {
		clear := buffer[i] ^ key
		if clear < ' ' || clear > '~' {
			errorCount++
		}
	}
	return errorCount
}

func printXor(buffer []byte, key byte) {
	for i := range buffer {
		clear := buffer[i] ^ key
		if clear < ' ' || clear > '~' {
			clear = '_'
		}
		fmt.Printf("%c", clear)
	}
	fmt.Println()
}

type Vector struct {
	vector       []int
	sumOfSquares float64
}

func createVector(buffer []byte, key byte) *Vector {

	var encoded Vector
	encoded.vector = make([]int, 256)

	for _, x := range buffer {
		encoded.vector[x^key]++
	}

	for _, x := range encoded.vector {
		encoded.sumOfSquares += float64(x * x)
	}

	return &encoded
}

func vectorAngle(vector1, vector2 *Vector) float64 {
	var dotProduct float64

	if len(vector1.vector) != len(vector2.vector) {
		log.Fatalf("Vectors not of same dimension: %d != %d\n", len(vector1.vector), len(vector2.vector))
	}

	for i, v1 := range vector1.vector {
		dotProduct += float64(v1 * vector2.vector[i])
	}

	magA := math.Sqrt(vector1.sumOfSquares)
	magB := math.Sqrt(vector2.sumOfSquares)
	z := dotProduct / (magA * magB)

	// math.Acos() undefined for argument -1 <= x >= 1,
	// and we know that z is positive.
	return math.Acos(z - 0.0000001)
}

var englishVector = &Vector{
	vector:       englishArray,
	sumOfSquares: englishSumOfSquares,
}

var englishArray = []int{
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	209,
	2989,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	17105,
	17,
	276,
	53,
	0,
	3,
	13,
	59,
	353,
	370,
	235,
	26,
	884,
	0,
	1561,
	793,
	310,
	228,
	187,
	92,
	61,
	51,
	66,
	59,
	82,
	73,
	384,
	10,
	75,
	402,
	86,
	7,
	8,
	348,
	177,
	344,
	246,
	323,
	97,
	296,
	86,
	480,
	85,
	12,
	175,
	143,
	238,
	207,
	292,
	4,
	286,
	393,
	425,
	149,
	54,
	77,
	88,
	55,
	6,
	307,
	9,
	309,
	1,
	304,
	0,
	5771,
	1861,
	3353,
	2887,
	9982,
	1769,
	2140,
	2663,
	6190,
	290,
	494,
	4378,
	2147,
	5410,
	6203,
	2120,
	78,
	5312,
	5498,
	7266,
	2649,
	999,
	928,
	292,
	1106,
	90,
	22,
	0,
	22,
	3,
	0,
	0,
	0,
	0,
	0,
	3,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	3,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	3,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
}
var englishSum float64 = 115075.0
var englishSumOfSquares float64 = 734987429.0
