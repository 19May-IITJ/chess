package board

import (
	"chess/modules/piece"
	"chess/utility"
	"fmt"
)

// Define a chessboard
type ChessBoard struct {
	Squares [8][8]*piece.ChessPiece
}

// Initialize a chessboard, factory function for chessboard
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{}
	// Set up the starting position (simplified for demonstration)
	board.Squares[0] = [8]*piece.ChessPiece{
		piece.NewChessPiece(utility.Rook, utility.Black, 0, 7),
		piece.NewChessPiece(utility.Knight, utility.Black, 0, 6),
		piece.NewChessPiece(utility.Bishop, utility.Black, 0, 5),
		piece.NewChessPiece(utility.Queen, utility.Black, 0, 4),
		piece.NewChessPiece(utility.King, utility.Black, 0, 3),
		piece.NewChessPiece(utility.Bishop, utility.Black, 0, 2),
		piece.NewChessPiece(utility.Knight, utility.Black, 0, 1),
		piece.NewChessPiece(utility.Rook, utility.Black, 0, 0),
	}
	board.Squares[1] = [8]*piece.ChessPiece{
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 7),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 6),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 5),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 4),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 3),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 2),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 1),
		piece.NewChessPiece(utility.Pawn, utility.Black, 0, 0),
	}

	board.Squares[7] = [8]*piece.ChessPiece{
		piece.NewChessPiece(utility.Rook, utility.White, 0, 7),
		piece.NewChessPiece(utility.Knight, utility.White, 0, 6),
		piece.NewChessPiece(utility.Bishop, utility.White, 0, 5),
		piece.NewChessPiece(utility.Queen, utility.White, 0, 4),
		piece.NewChessPiece(utility.King, utility.White, 0, 3),
		piece.NewChessPiece(utility.Bishop, utility.White, 0, 2),
		piece.NewChessPiece(utility.Knight, utility.White, 0, 1),
		piece.NewChessPiece(utility.Rook, utility.White, 0, 0),
	}
	board.Squares[6] = [8]*piece.ChessPiece{
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 7),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 6),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 5),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 4),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 3),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 2),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 1),
		piece.NewChessPiece(utility.Pawn, utility.White, 0, 0),
	}

	return board
}

func (board *ChessBoard) PrintBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%d ", board.Squares[i][j])
		}
		fmt.Println()
	}
}
