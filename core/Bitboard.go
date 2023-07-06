package core

import (
	"fmt"
	"strings"
)

type Bitboard struct {
	whitePawns   uint64
	blackPawns   uint64
	whiteKnights uint64
	blackKnights uint64
	whiteBishops uint64
	blackBishops uint64
	whiteRooks   uint64
	blackRooks   uint64
	whiteQueen   uint64
	blackQueen   uint64
	whiteKing    uint64
	blackKing    uint64
}

func NewChessboard() *Bitboard {
	bb := &Bitboard{}
	bb.whitePawns = 0x000000000000FF00
	bb.whiteKnights = 0x0000000000000042
	bb.whiteBishops = 0x0000000000000024
	bb.whiteRooks = 0x0000000000000081
	bb.whiteQueen = 0x0000000000000008
	bb.whiteKing = 0x0000000000000010
	bb.blackPawns = 0x00FF000000000000
	bb.blackKnights = 0x4200000000000000
	bb.blackBishops = 0x2400000000000000
	bb.blackRooks = 0x8100000000000000
	bb.blackQueen = 0x0800000000000000
	bb.blackKing = 0x1000000000000000
	return bb
}

func FENtoBitboard(fen string) *Bitboard {
	bb := new(Bitboard)
	// Split the FEN string into its components
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")

	// Iterate over the ranks and files to populate the bitboard structure
	for rank, rankStr := range ranks {
		file := 0
		for _, char := range rankStr {
			if char >= '1' && char <= '8' {
				file += int(char - '0')
			} else {
				index := (8 * (7 - rank)) + file
				switch char {
				case 'P':
					bb.whitePawns |= (1 << index)
				case 'N':
					bb.whiteKnights |= (1 << index)
				case 'B':
					bb.whiteBishops |= (1 << index)
				case 'R':
					bb.whiteRooks |= (1 << index)
				case 'Q':
					bb.whiteQueen |= (1 << index)
				case 'K':
					bb.whiteKing |= (1 << index)
				case 'p':
					bb.blackPawns |= (1 << index)
				case 'n':
					bb.blackKnights |= (1 << index)
				case 'b':
					bb.blackBishops |= (1 << index)
				case 'r':
					bb.blackRooks |= (1 << index)
				case 'q':
					bb.blackQueen |= (1 << index)
				case 'k':
					bb.blackKing |= (1 << index)
				}
				file++
			}
		}
	}

	return bb
}

func (board *Bitboard) Display() {
	// Print top border
	fmt.Println("  +------------------------+")

	// Loop through ranks (rows) in reverse order
	for rank := 7; rank >= 0; rank-- {
		fmt.Printf("%d |", rank+1)

		// Loop through files (columns)
		for file := 0; file < 8; file++ {
			square := uint64(1) << (rank*8 + file)
			var piece string

			if board.whitePawns&square != 0 {
				piece = "P"
			} else if board.blackPawns&square != 0 {
				piece = "p"
			} else if board.whiteKnights&square != 0 {
				piece = "N"
			} else if board.blackKnights&square != 0 {
				piece = "n"
			} else if board.whiteBishops&square != 0 {
				piece = "B"
			} else if board.blackBishops&square != 0 {
				piece = "b"
			} else if board.whiteRooks&square != 0 {
				piece = "R"
			} else if board.blackRooks&square != 0 {
				piece = "r"
			} else if board.whiteQueen&square != 0 {
				piece = "Q"
			} else if board.blackQueen&square != 0 {
				piece = "q"
			} else if board.whiteKing&square != 0 {
				piece = "K"
			} else if board.blackKing&square != 0 {
				piece = "k"
			} else {
				piece = "*"
			}

			fmt.Printf(" %s", piece)
		}

		// Print right border and rank number
		fmt.Printf(" | %d\n", rank+1)
	}

	// Print bottom border
	fmt.Println("  +------------------------+")

	// Print file letters
	fmt.Println("    a b c d e f g h")
}
func (bb *Bitboard) DisplayDebug() {

	fmt.Printf("Pions blancs : %064b\n", bb.whitePawns)
	fmt.Printf("Cavaliers blancs : %064b\n", bb.whiteKnights)
	fmt.Printf("Fous blancs : %064b\n", bb.whiteBishops)
	fmt.Printf("Tours blanches : %064b\n", bb.whiteRooks)
	fmt.Printf("Reine blanche : %064b\n", bb.whiteQueen)
	fmt.Printf("Roi blanc : %064b\n", bb.whiteKing)

	// Affichage des bitboards pour chaque pièce des noirs
	fmt.Printf("Pions noirs : %064b\n", bb.blackPawns)
	fmt.Printf("Cavaliers noirs : %064b\n", bb.blackKnights)
	fmt.Printf("Fous noirs : %064b\n", bb.blackBishops)
	fmt.Printf("Tours noires : %064b\n", bb.blackRooks)
	fmt.Printf("Reine noire : %064b\n", bb.blackQueen)
	fmt.Printf("Roi noir : %064b\n", bb.blackKing)
}
