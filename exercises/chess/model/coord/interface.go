// Package coord will contain types and logic for handling chess board coordinates.
package coord

import "fmt"

// ChessCoordinates is an interface for an abstract Coordinates.
type ChessCoordinates interface {
	fmt.Stringer // "A7"
	// Coord gets n'th coordinate comp.
	// Start with 0th coordinate.
	// Returns an error if coordinate is inexistant.
	Coord(n int) (int, error)

	// Vector
	// Add()
	// Multiply()
}
