package piece

// Define a struct for a chess piece
type ChessPiece struct {
	Type  int // Type of the piece (Pawn, Knight, etc.)
	Color int // Color of the piece (White or Black)
	Row   int // Row position on the chessboard
	Col   int // Column position on the chessboard
}

func NewChessPiece(type_, color, row, col int) *ChessPiece {
	return &ChessPiece{
		Type:  type_,
		Color: color,
		Row:   row,
		Col:   col,
	}
}
