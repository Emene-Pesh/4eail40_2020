package main

import (
	"fmt"
	"io"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type MemoryStorage struct{}

func (m MemoryStorage) Open(s string) (io.ReadWriteCloser, error) {
	fmt.Printf("Opening %s\n", s)
	return nil, nil
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

type DiskStorage struct{}

func (d DiskStorage) Open(s string) (io.ReadWriteCloser, error) {
	fmt.Printf("Opening %s\n", s)
	return nil, nil
}

func NewDiskStorage() *DiskStorage {
	return &DiskStorage{}
}

type TempStorage struct{}

func (t TempStorage) Open(s string) (io.ReadWriteCloser, error) {
	fmt.Printf("Opening %s\n", s)
	return nil, nil
}

func NewTempStorage() *TempStorage {
	return &TempStorage{}
}
