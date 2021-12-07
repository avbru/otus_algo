package main

import (
	"fmt"
	"log"
	"os"
)

const usage = "usage: -encode || -decode -rle || -rlePlus in.file out.file"

func main() {
	args := os.Args
	fmt.Println(args)
	if len(args) != 5 {
		log.Fatalln(usage)
	}

	var Encode func([]byte) []byte
	var Decode func([]byte) []byte

	switch args[2] {
	case "-rle":
		Encode = encode
		Decode = decode
	case "-rlePlus":
		Encode = encodeP
		Decode = decodeP
	default:
		log.Fatalln(usage)
	}

	in, err := os.ReadFile(args[3])
	if err != nil {
		log.Fatalln(err)
	}

	switch args[1] {
	case "-encode":
		if err := os.WriteFile(os.Args[4], Encode(in), 0666); err != nil {
			log.Fatalln(err)
		}
	case "-decode":
		if err := os.WriteFile(os.Args[4], Decode(in), 0666); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln(usage)
	}

}
