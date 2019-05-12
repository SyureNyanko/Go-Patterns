package main

import (
	"fmt"
)

type Store interface {
	Open(string) error
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

/* factory pattern */
func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()

	}
}

/* This memoryStore struct is qualified Store interface,
so treat as Store interface */
type memoryStore struct{}

func (ms *memoryStore) Open(p string) error {
	fmt.Printf("memory store Open %s \n", p)
	return nil
}

func newMemoryStorage() Store {
	return &memoryStore{}
}

type diskStore struct{}

func (ms *diskStore) Open(p string) error {
	fmt.Printf("disk store Open %s \n", p)
	return nil
}

func newDiskStorage() Store {
	return &diskStore{}
}

type tempStore struct{}

func (ms *tempStore) Open(p string) error {
	fmt.Printf("temp store Open %s \n", p)
	return nil
}

func newTempStorage() Store {
	return &tempStore{}
}

func main() {
	s := NewStore(MemoryStorage)
	s.Open("file")
}
