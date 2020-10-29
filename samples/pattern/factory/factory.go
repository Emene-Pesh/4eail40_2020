package main

import "errors"

type StorageType int

const (
	Disk StorageType = 1 << iota
	Tmp
	Memory
)

func NewStore(t StorageType) (Store, error) {
	switch t {
	case Memory:
		return NewMemoryStorage( /*...*/ ), nil
	case Disk:
		return NewDiskStorage( /*...*/ ), nil
	case Tmp:
		return NewTempStorage( /*...*/ ), nil
	}
	return nil, errors.New("unknown storage type")
}
