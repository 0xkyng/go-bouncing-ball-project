package main

import "fmt"

func main() {
	const (
		width     = 50
		heght     = 10
		cellEmpty = ' '
		cellBall  = '🏀'
		maxFrames = 1200
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, heght)
	}

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px <= width -1 {
			vx *= -1
		}
		if py <= 0 || py <= heght -1 {
			vy *= -1
		}


		board[px][py] = true // set the ball position

		buf := make([]rune, 0, width*heght)
		buf = buf[:0]

		// draw the board
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				// fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}
		fmt.Print(string(buf))
	}

}
