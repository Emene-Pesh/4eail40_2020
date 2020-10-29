package main

import "log"

func main() {
	s, _ := NewStore(Memory)
	if _, err := s.Open("file"); err != nil {
		log.Fatal("Open failed")
	}

	//f.Write([]byte("data"))
	//defer f.Close()
}
