package solver

import (
	"testing"
)

var solvedValidSudoku = []int{
	8, 6, 4, 3, 7, 1, 2, 5, 9,
	3, 2, 5, 8, 4, 9, 7, 6, 1,
	9, 7, 1, 2, 6, 5, 8, 4, 3,
	4, 3, 6, 1, 9, 2, 5, 8, 7,
	1, 9, 8, 6, 5, 7, 4, 3, 2,
	2, 5, 7, 4, 8, 3, 9, 1, 6,
	6, 8, 9, 7, 3, 4, 1, 2, 5,
	7, 1, 3, 5, 2, 8, 6, 9, 4,
	5, 4, 2, 9, 1, 6, 3, 7, 8,
}
var solvedInvalidSudoku = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	1, 2, 3, 4, 5, 6, 7, 8, 9,
}
var notSolvedValidSudoku = []int{
	0, 0, 4, 3, 0, 0, 2, 0, 9,
	0, 0, 5, 0, 0, 9, 0, 0, 1,
	0, 7, 0, 0, 6, 0, 0, 4, 3,
	0, 0, 6, 0, 0, 2, 0, 8, 7,
	1, 9, 0, 0, 0, 7, 4, 0, 0,
	0, 5, 0, 0, 8, 3, 0, 0, 0,
	6, 0, 0, 0, 0, 0, 1, 0, 5,
	0, 0, 3, 5, 0, 8, 6, 9, 0,
	0, 4, 2, 9, 1, 0, 3, 0, 0,
}

func Test_isValid(t *testing.T) {

	t.Run("not valid sudoku", func(t *testing.T) {
		valid := isValid(solvedInvalidSudoku)

		if valid {
			t.Errorf("should have failed, it is not a valid sudoku. want false, got %t", valid)
		}
	})

	t.Run("valid unsolved sudoku", func(t *testing.T) {
		notValid := !isValid(solvedValidSudoku)

		if notValid {
			t.Errorf("should have passed, it is a valid sudoku. want true, got %t", notValid)
		}
	})

	t.Run("valid solved sudoku", func(t *testing.T) {
		notValid := !isValid(notSolvedValidSudoku)

		if notValid {
			t.Errorf("should have passed, it is a valid sudoku. want true, got %t", notValid)
		}
	})
}

func Test_isSolved(t *testing.T) {

	t.Run("test is solved util method for not solved puzzle", func(t *testing.T) {
		solved := isSolved(notSolvedValidSudoku)

		if solved {
			t.Errorf("it is a not solved sudoku, but isSolved returned %t", solved)
		}
	})

	t.Run("test is solved util method for solved puzzle", func(t *testing.T) {
		solved := isSolved(solvedValidSudoku)

		if !solved {
			t.Errorf("it is a solved sudoku, but isSolved returned %t", solved)
		}
	})
}

func Test_findNextIndexToSolve(t *testing.T) {

	t.Run("test finding the next index to solve for a not solved puzzle", func(t *testing.T) {
		nextIndex := findNextIndexToSolve(notSolvedValidSudoku)

		if nextIndex != 0 {
			t.Errorf("the index to find next should be 0, but findNextIndexToSolve returned %d", nextIndex)
		}
	})

	t.Run("test finding the next index to solve for a solved puzzle", func(t *testing.T) {

		nextIndex := findNextIndexToSolve(solvedValidSudoku)

		if nextIndex != -1 {
			t.Errorf("the index to find next should be -1, but findNextIndexToSolve returned %d", nextIndex)
		}
	})
}
