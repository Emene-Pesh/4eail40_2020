package main

import "fmt"

func main() {
	s := NewSingleton()

	s["this"] = "that"

	s2 := NewSingleton()

	fmt.Println("This is ", s2["this"])
	// This is that
}
