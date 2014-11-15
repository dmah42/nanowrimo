package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestNewChain(t *testing.T) {
	chain := map[string][]string {
		" ": []string{"i"},
		" i": []string{"think"},
		"i think": []string{"therefore"},
		"think therefore": []string{"i"},
		"therefore i": []string{"am"},
	}
	c := NewChain(2, strings.NewReader("I think therefore I am"))
	if c.prefixLen != 2 {
		t.Error("Expected 2, got ", c.prefixLen)
	}

	if !reflect.DeepEqual(c.chain, chain) {
		got, _ := json.Marshal(c.chain)
		want, _ := json.Marshal(chain)
		t.Error(fmt.Printf("Expected %s\ngot      %s", want, got))
	}
}
