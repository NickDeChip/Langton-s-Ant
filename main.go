package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	winWidth  = 800
	winHeight = 800
	scl       = 5
	rows      = winWidth / scl
	cols      = winHeight / scl
)

const (
	up = iota
	left
	down
	right
)

var (
	grid [rows][cols]int

	cycles int

	antX   = rows / 2
	antY   = cols / 2
	antDir = left
)

func main() {
	rl.InitWindow(winWidth, winHeight, "Ant!?")
	rl.SetTargetFPS(0)

	for !rl.WindowShouldClose() {

		for i := 0; i < 10; i++ {
			cycles += 1
			antMove()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		draw()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func draw() {
	for i := range grid {
		for j := range grid[i] {
			state := grid[i][j]
			switch state {
			case 0:
				rl.DrawRectangle(int32(i)*scl, int32(j)*scl, scl, scl, rl.White)
			case 1:
				rl.DrawRectangle(int32(i)*scl, int32(j)*scl, scl, scl, rl.Black)
			}
			rl.DrawText(fmt.Sprintf("CYCLES: %d", cycles), 10, 10, 32, rl.Red)
			rl.DrawRectangle(int32(antX)*scl, int32(antY)*scl, scl, scl, rl.Red)
		}
	}
}

func antMove() {
	cell := grid[antX][antY]
	if cell == 0 {
		grid[antX][antY] = 1
		switch antDir {
		case up:
			antDir = right
			antX++
		case left:
			antDir = up
			antY--
		case down:
			antDir = left
			antX--
		case right:
			antDir = down
			antY++
		}
	} else {
		grid[antX][antY] = 0
		switch antDir {
		case up:
			antDir = left
			antX--
		case left:
			antDir = down
			antY++
		case down:
			antDir = right
			antX++
		case right:
			antDir = up
			antY--
		}
	}

	if antX < 0 {
		antX = rows - 1
	} else if antX >= rows {
		antX = 0
	}

	if antY < 0 {
		antY = cols - 1
	} else if antY >= cols {
		antY = 0
	}
}
