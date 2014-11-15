package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	numStories := flag.Int("stories", 1, "number of stories to generate")
	numWords := flag.Int("words", 100, "maximum number of words per story")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	c := NewChain(*prefixLen, os.Stdin)
	for i := 0; i < *numStories; i++ {
		fmt.Printf("%d. %s\n", i, c.Generate(*numWords))
	}
}
