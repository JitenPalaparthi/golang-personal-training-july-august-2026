package main

import (
	"fmt"
	"hash/fnv"
)

func main() {

	key := "560086"
	hash := GetHash64(key)
	fmt.Printf("Key:%s\n", key)
	fmt.Printf("Hash:%d\n", hash)
	fmt.Printf("Hex:0x%x\n", hash)
	fmt.Printf("Hex:%064b\n", hash)

	// 001000010010101001001000010000111001110100010110001101011 1001000
	// h1 ------------------------------------------------------   h2
	// &
	// 000000000000000000000000000000000000000000000000000000000 1111111
	// 000000000000000000000000000000000000000000000000000000000 1001000

	// 0000000 001000010010101001001000010000111001110100010110001101011
	// h1 and h2

	//fmt.Println(0x7f)
	fmt.Printf("%064b\n", 0x7f) // This is only the number which has 7 bits as 1111111

	h2 := hash & 0b1111111
	fmt.Println(h2)
	fmt.Printf("%07b\n", h2)

	h1 := hash >> 7
	fmt.Printf("%064b\n", h1)

	fmt.Println("mod(assume bucket):", h1%4) // arthimetic mod
	fmt.Println("mod(assume bucket index):", h2%8)
	fmt.Printf("%08b %08b %08b %08b %08b\n", 2-1, 4-1, 8-1, 16-1, 32-1) // 2 powerof n -1
	bktIndex := h1 & (4 - 1)                                            // bitwise mod
	fmt.Println("bucket index:", bktIndex)
	indexBct := h2 & (8 - 1)
	fmt.Println("index inside bucket:", indexBct)
	println(128 * (7.0 / 8.0))

}

func GetHash64(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

// 2
// 4
// 8
// 16
