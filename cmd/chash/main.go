package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func sha256sum(d []byte) []byte {
	h := sha256.New()
	h.Write(d)
	return h.Sum(nil)
}

func main() {
	f := ""
	if len(os.Args) == 2 {
		f = os.Args[1]
	}

	contents, err := os.ReadFile(f)

	if err != nil {
		fmt.Println("Error opening file", f, err)
		os.Exit(1)
	}

	sum := sha256sum(contents)

	fmt.Println(fmt.Sprintf("%x", sum))
}
