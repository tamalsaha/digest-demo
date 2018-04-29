package main

import (
	"github.com/opencontainers/go-digest"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"log"
	"fmt"
)

func main() {
	d, err := digest.Parse("sha256:7173b809ca12ec5dee4506cd86be934c4596dd234ee82c0662eac04a8c2c71dc")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(d.Algorithm())
	fmt.Println(d.String())
	fmt.Println(d.Encoded())
	fmt.Println(d.Hex())
}
