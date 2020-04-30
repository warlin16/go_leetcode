package main

import (
	"math"
	"strings"
)

/**
On an alphabet board, we start at position (0, 0),
corresponding to character board[0][0].

Here,
board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"],
as shown in the diagram below.

'U' moves our position up one row, if the position exists on the board;
'D' moves our position down one row, if the position exists on the board;
'L' moves our position left one column, if the position exists on the board;
'R' moves our position right one column, if the position exists on the board;
'!' adds the character board[r][c] at our current position (r, c) to the answer.
(Here, the only positions that exist on the board are positions with letters on them.)

Return a sequence of moves that makes our answer equal to target in the minimum number of
moves.  You may return any path that does so.

Example 1:

Input: target = "leet"
Output: "DDR!UURRR!!DDD!"
Example 2:

Input: target = "code"
Output: "RR!DDRR!UUL!R!"
*/

var board = []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}

func alphabetBoardPath(target string) string {
	sequence := ""
	alphabet := strings.Join(board, "")
	coord := [2]int{0, 0}
	for i := range target {
		letter := string(target[i])
		idxInAlphabet := strings.Index(alphabet, letter)
		coord = calculateMovement(coord, idxInAlphabet, letter, &sequence)
	}
	return sequence
}

func calculateMovement(coord [2]int, index int, letter string, sequence *string) [2]int {
	if coord == [2]int{5, 0} && letter == "z" {
		*sequence += "!"
		return coord
	}
	xCoord := int(math.Floor(float64(index) / 5))
	yCoord := strings.Index(board[int(xCoord)], letter)
	xTravel := xCoord - coord[0]
	if letter == "z" {
		xTravel--
	}
	if xTravel < 0 {
		*sequence += strings.Repeat("U", -xTravel)
	} else if xTravel > 0 {
		*sequence += strings.Repeat("D", xTravel)
	}
	yTravel := yCoord - coord[1]
	if yTravel < 0 {
		*sequence += strings.Repeat("L", -yTravel)
	} else if yTravel > 0 {
		*sequence += strings.Repeat("R", yTravel)
	}
	if letter == "z" {
		*sequence += "D"
	}
	*sequence += "!"
	return [2]int{xCoord, yCoord}
}
