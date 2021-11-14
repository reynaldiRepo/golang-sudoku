package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//check sudoku board is valid
func isBoardValid(board [][]int) bool {
	if len(board) != 9 {
		return false
	}
	for i := 0; i < len(board); i++ {
		if len(board[i]) != 9 {
			return false
		}
	}
	return true
}

//check if sudoku board is solved
func isSolved(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

//check is valid value
func isValidValue(board [][]int, row int, col int, value int) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] == value {
			return false
		}
	}
	for i := 0; i < len(board[row]); i++ {
		if board[i][col] == value {
			return false
		}
	}
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == value {
				return false
			}
		}
	}
	return true
}

func findEmptyCell(board [][]int) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return 0, 0
}

// function to solve sudoku
func solveSudoku(board [][]int) ([][]int, bool, string) {

	// check if board is valid
	if !isBoardValid(board) {
		return board, false, "board is not valid"
	}

	// check if board is solved
	if isSolved(board) {
		fmt.Println("Sudoku is solved", board)
		return board, true, "sudoku is solved"
	}

	// find the first empty cell
	row, col := findEmptyCell(board)

	// loop through all possible values
	for i := 1; i <= 9; i++ {
		// check if value is valid
		if isValidValue(board, row, col, i) {
			// set value
			board[row][col] = i
			// solve sudoku
			board, _, _ = solveSudoku(board)
			if isSolved(board) {
				return board, true, "sudoku is solved"
			}
			// reset value
			board[row][col] = 0
		}
	}

	// check if board is solved after filling
	if isSolved(board) {
		fmt.Println("Sudoku is solved on outside loop", board)
		return board, true, "sudoku is solved"
	} else {
		return board, false, "sudoku is not solved"
	}
}

type INPUT struct {
	Input [][]int `binding:"required"`
}

func handleRequestSudoku(c *gin.Context) {
	data := new(INPUT)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid input",
			"err":     err.Error(),
		})
		return
	}
	fmt.Println("input data", data.Input)
	SolvedSudoku, isSolved, message := solveSudoku(data.Input)
	c.JSON(200, gin.H{
		"IsSolved": isSolved,
		"message":  message,
		"sudoku":   SolvedSudoku,
	})
}

func main() {

	//==============test run sudoku function on first run===================
	board := [][]int{
		{0, 3, 0, 0, 0, 0, 8, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 4, 2},
		{2, 0, 8, 6, 7, 0, 3, 0, 5},
		{8, 5, 0, 0, 1, 0, 6, 2, 0},
		{0, 0, 7, 0, 0, 0, 9, 0, 0},
		{0, 4, 9, 0, 5, 0, 0, 1, 8},
		{9, 0, 5, 0, 4, 7, 2, 0, 6},
		{3, 7, 0, 0, 0, 6, 4, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 7, 0},
	}
	fmt.Println("Testing Solvig sudoku...", board)
	SolvedSudoku, _, _ := solveSudoku(board)
	fmt.Println("Testing Solved sudoku...", SolvedSudoku)
	//============= test run sudoku function on run

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"massage": "Reynaldi Backend Developer Test",
			"Task":    "Create API for solving sudoku problem @idn_media",
			"Note":    "API endpoint for solving sudoku at /sudoku, input should be json request of array with size 9 x9",
		})
	})

	r.POST("/sudoku", handleRequestSudoku)

	r.Run()
}
