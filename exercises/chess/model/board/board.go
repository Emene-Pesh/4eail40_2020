package board

import (
	"fmt"
	"github.com/Emene-Pesh/4eail40_2020/exercises/chess/model/coord"
	"github.com/Emene-Pesh/4eail40_2020/exercises/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	x,_ := at.Coord(0)
	y,_ := at.Coord(1)

	if x>8 || y>8 {
		return nil
	}
	piece := c[x][y]
	return piece
}

// MovePiece moves a piece from given coordinates to
// given coordinates.
// Returns an error if destination was occupied.
func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	if c.PieceAt(to) != nil{
		return fmt.Errorf("occupied")
	}
	if c.PieceAt(from) == nil{
		return fmt.Errorf("No pieces to move" )
	}
	x,_ := from.Coord(0)
	y,_ := from.Coord(1)
	piece := c[x][y]
	c[x][y]=nil
	return c.PlacePieceAt(piece,to)

}

// PlacePieceAt places a given piece at given location.
// Returns an error if destination was occupied.
func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	if c.PieceAt(at) != nil{
		return fmt.Errorf("occupied")
	}
	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	c[x][y] = p
	return nil
}
