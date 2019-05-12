package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

func main() {
	s := New()
	s["I"] = "syurenyanko"
	s2 := New()

	fmt.Println("I am ", s2["I"])
}
