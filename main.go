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
		if board[row][i] == value && i != col {
			return false
		}
	}
	for i := 0; i < len(board[row]); i++ {
		if board[i][col] == value && i != row {
			return false
		}
	}

	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == value && i != row && j != col {
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
			"is_solved" : false,
			"message": err.Error(),
			"note" : "input should be json request of array with size 9 x 9 inside 'Input' key",
		})
		return
	}
	fmt.Println("input data", data.Input)

	//check if input valid sudoku board
	if !isBoardValid(data.Input) {
		c.JSON(200, gin.H{
			"is_solved": false,
			"message":  "Board size is invalid",
			"sudoku":   data.Input,
		})
		return
	}

	//check if input has invalid values
	for i := 0; i < len(data.Input); i++ {
		for j := 0; j < len(data.Input[i]); j++ {
			if !isValidValue(data.Input, i, j, data.Input[i][j]) && data.Input[i][j] != 0 {
				fmt.Println(isValidValue(data.Input, i, j, data.Input[i][j]))
				c.JSON(200, gin.H{
					"is_solved": false,
					"message":  fmt.Sprintf("value %d is invalid at (row %d , col %d )", data.Input[i][j], i, j),
					"sudoku":   data.Input,
				})
				return
			}
		}
	}

	SolvedSudoku, isSolved, message := solveSudoku(data.Input)
	c.JSON(200, gin.H{
		"is_solved": isSolved,
		"message":  message,
		"sudoku":   SolvedSudoku,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"massage": "Reynaldi Backend Developer Test",
			"Task":    "Create API for solving sudoku problem @idn_media",
			"Note":    "API endpoint for solving sudoku at /sudoku, input should be json request of array with size 9 x 9",
		})
	})

	r.POST("/sudoku", handleRequestSudoku)

	r.Run()
}
