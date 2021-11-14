package main

import "testing"

func TestValidationSizeBoard(t *testing.T) {
	InvalidBoard := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	if isBoardValid(InvalidBoard) {
		t.Error("Invalid size of board should not be valid")
	} else {
		t.Log("Invalid size board is not valid")
	}

	validBoard := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	if isBoardValid(validBoard) {
		t.Log("Valid board is valid")
	} else {
		t.Error("Valid board should be valid")
	}

}

//create test for valid value on function isValidValue
func TestFindEmptyCell(t *testing.T) {
	board := [][]int{
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 8, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 5, 0, 3, 6},
		{3, 0, 4, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 2, 0},
	}
	row, col := findEmptyCell(board)
	if row != 0 || col != 0 {
		t.Error("should be on row 0 col 0")
	} else {
		t.Log("Find empty cell is correct")
	}

	board = [][]int{
		{1, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 8, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 5, 0, 3, 6},
		{3, 0, 4, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 2, 0},
	}
	row, col = findEmptyCell(board)
	if row != 0 || col != 2 {
		t.Error("should be on row 0 col 0")
	} else {
		t.Log("Find empty cell is correct")
	}
}

func TestIsValidValue(t *testing.T) {
	board := [][]int{
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 8, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 5, 0, 3, 6},
		{3, 0, 4, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 2, 0},
	}

	if isValidValue(board, 0, 0, 2) {
		t.Error("should be invalid")
	} else {
		t.Log("Invalid value is invalid")
	}

	if isValidValue(board, 0, 0, 1) {
		t.Log("Valid Valie")
	} else {
		t.Error("Should be valid")
	}
}

func TestSudokuSolver(t *testing.T) {
	board := [][]int{
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 8, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 5, 0, 3, 6},
		{3, 0, 4, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 2, 0},
	}
	_, status, _ := solveSudoku(board)
	if status != true {
		t.Error("Should be valid")
	} else {
		t.Log("Sudoku is solved")
	}

	board = [][]int{
		{1, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 8, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 5, 0, 3, 6},
		{3, 0, 4, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 2, 0},
	}
	_, status, _ = solveSudoku(board)
	if status != true {
		t.Error("Should be valid")
	} else {
		t.Log("Sudoku is solved")
	}

	board = [][]int{
		{1, 2, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 0, 0, 0, 7, 0, 8, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{0, 0, 0, 0, 0, 5, 0, 3, 6},
		{3, 0, 4, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 9},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 0, 0, 0, 2, 0},
	}
	_, status, _ = solveSudoku(board)
	if status != true {
		t.Error("Should be valid")
	} else {
		t.Log("Sudoku is solved")
	}
}
