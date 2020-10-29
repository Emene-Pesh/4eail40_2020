// Package board will contain types and logic for handling chess board(s).
package board

import (
	"fmt"

	"github.com/jglouis/4eail40_2020/exercises/chess/model/coord"
	"github.com/jglouis/4eail40_2020/exercises/chess/model/piece"
)

// Board is an interface to a chess board.
// It defines methods for handling pieces on it.
type Board interface {
	fmt.Stringer
	// PieceAt retrievves piece at give coordinates.
	// Returns nil if no piece was found.
	PieceAt(at coord.ChessCoordinates) piece.Piece
	// MovePiece moves a piece from given coordinates to
	// given coordinates.
	// Returns an error if destination was occupied.
	MovePiece(from, to coord.ChessCoordinates) error
	// PlacePieceAt places a given piece at given location.
	// Returns an error if destination was occupied.
	PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error
}
