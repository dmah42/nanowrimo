package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type prefix []string

func (p prefix) String() string {
	return strings.Join(p, " ")
}

func (p prefix) shift(word string) {
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
	p := make(prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		// Ensure strings are in lower case.
		s = strings.ToLower(s)

		// TODO: strip off any punctuation and add the punctuation as a potential suffix.
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.shift(s)
	}
	return c
}

func (c *Chain) Generate(n int) string {
	p := make(prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		// TODO: capitalise after period, etc.
		p.shift(next)
	}
	return strings.Join(words, " ")
}

