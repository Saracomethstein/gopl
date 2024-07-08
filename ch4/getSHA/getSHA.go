package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var t = flag.Bool("t", false, "sha384")
var f = flag.Bool("f", false, "sha512")

func main() {
	parse()
}

func getSHA256(data string) [32]byte {
	result := sha256.Sum256([]byte(data))
	return result
}

func getSHA384(data string) [48]byte {
	result := sha512.Sum384([]byte(data))
	return result
}

func getSHA512(data string) [64]byte {
	result := sha512.Sum512([]byte(data))
	return result
}

func parse() {
	flag.Parse()
	massage := os.Args[1:]
	if *t {
		fmt.Printf("%s\n%x\n", massage[1], getSHA384(massage[1]))
	} else if *f {
		fmt.Printf("%s\n%x\n", massage[1], getSHA512(massage[1]))
	} else {
		fmt.Printf("%s\n%x\n", massage[0], getSHA256(massage[0]))
	}
}
