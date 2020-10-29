package main

import "log"

type IntToInt func(int) int

func LogDecorate(fn IntToInt) IntToInt {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
