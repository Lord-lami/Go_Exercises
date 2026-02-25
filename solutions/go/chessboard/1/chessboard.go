package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

// File type stores the vacancy of vertical columns of chessboard.
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// Chessboard type stores the vacancy of all the cells of a chessboard.
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (count int) {
	if column, valid := cb[file]; valid {
		for _, occupied := range column {
			if occupied {
				count++
			}
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (count int) {
	if rank >= 1 && rank <= 8 {
		for _, occupied := range cb {
			if occupied[rank-1] {
				count++
			}
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (count int) {
	for range cb {
		count += 8
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (count int) {
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
	}
	return
}
