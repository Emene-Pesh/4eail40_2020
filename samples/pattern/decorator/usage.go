package main

func Double(n int) int {
	return n * 2
}

func main() {
	f := LogDecorate(Double)

	f(5)
	// Starting execution with the integer 5
	// Execution is completed with the result 10
}
