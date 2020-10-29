package board

import (
	"reflect"
	"testing"

	"github.com/Emene-Pesh/4eail40_2020/exercises/chess/model/coord"
	"github.com/Emene-Pesh/4eail40_2020/exercises/chess/model/piece"
	"github.com/Emene-Pesh/4eail40_2020/exercises/chess/model/player"
)

type mockCoord int

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

type mockPiece struct {
}

func (p mockPiece) String() string {
	panic(" ")
}
func (p mockPiece) Color() player.Color {
	panic("not implemented")
}
func (p mockPiece) Moves(b bool) map[coord.ChessCoordinates]bool {
	panic("")
}

func TestClassic_MovePiece(t *testing.T) {
	ca := Classic{}
	pto := mockPiece{}
	pfr := mockPiece{}
	coordto:= coord.NewCartesian(0,0)
	x, _ := coordto.Coord(0)
	y, _ := coordto.Coord(1)
	ca[x][y] = pto
	coordfr := coord.NewCartesian(7, 0)
	x1, _ := coordfr.Coord(0)
	y1, _ := coordfr.Coord(1)
	ca[x1][y1] = pfr

	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       *Classic
		args    args
		wantErr bool // true if there is an error, false if there is not.
	}{
		{
			"No piece",
			&ca,
			args{from: coord.NewCartesian(3, 5), to: coord.NewCartesian(3, 5)},
			true,
		},
		{
			"Occupied",
			&ca,
			args{from: coordfr, to: coordto},
			true,
		},
		{
			"Legal Move",
			&ca,
			args{from: coordfr, to: coord.NewCartesian(4, 4)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}



func TestClassic_PlacePieceAt(t *testing.T) {
	ca := Classic{}
	p1 := mockPiece{}
	place := coord.NewCartesian(0, 0)
	x, _ := place.Coord(0)
	y, _ := place.Coord(1)
	ca[x][y] = p1
	type args struct {
		p  piece.Piece
		at coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       *Classic
		args    args
		wantErr bool
	}{
		{
			"occupied",
			&ca,
			args{p1, place},
			true,
		},
		{
			"Empty",
			&ca,
			args{p1, coord.NewCartesian(2, 2)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.PlacePieceAt(tt.args.p, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("Classic.PlacePieceAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClassic_PieceAt(t *testing.T) {
	ca := Classic{}
	p1 := mockPiece{}
	place := coord.NewCartesian(0, 0)
	x, _ := place.Coord(0)
	y, _ := place.Coord(1)
	ca[x][y] = p1
	type args struct {
		at coord.ChessCoordinates
	}
	tests := []struct {
		name string
		c    Classic
		args args
		want piece.Piece
	}{
		{
			"There is",
			ca,
			args{at: place},
			p1,
		},
		{
			"No piece",
			ca,
			args{at: coord.NewCartesian(5, 5)},
			nil,
		},
		{
			"Out of bounds",
			ca,
			args{at: coord.NewCartesian(9, 9)},
			nil,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PieceAt(tt.args.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Classic.PieceAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
