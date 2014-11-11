package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Key []string

func (p Key) String() string {
	return strings.Join(p, " ")
}

func (p Key) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int, r io.Reader) *Chain {
	c := &Chain{make(map[string][]string), prefixLen}

	br := bufio.NewReader(r)
	k := make(Key, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		// Ensure strings are in lower case.
		s = strings.ToLower(s)

		// TODO: strip off any punctuation and add the punctuation as a potential suffix.
		key := k.String()
		c.chain[key] = append(c.chain[key], s)
		k.Shift(s)
	}
	return c
}

func (c *Chain) Generate(n int) string {
	k := make(Key, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[k.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		// TODO: capitalise after period, etc.
		k.Shift(next)
	}
	return strings.Join(words, " ")
}

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
