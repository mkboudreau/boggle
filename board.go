package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BoardSize int8

type Board [][]byte

type Coord struct {
	x, y int
}

const (
	Board3x3 BoardSize = 3
	Board4x4           = 4
	Board5x5           = 5
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZAEIOUAEIOUSTRNL"

var random = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

func NewBoard(size BoardSize) Board {
	board := make([][]byte, size)
	for i, _ := range board {
		board[i] = make([]byte, size)
	}
	for x := 0; x < int(size); x++ {
		for y := 0; y < int(size); y++ {
			board[x][y] = randomLetter()
		}
	}
	return board
}

func randomLetter() byte {
	index := random.Intn(len(letters))
	return letters[index]
}

func (b Board) String() string {
	var s string
	rows := len(b)
	cols := len(b[0])

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			s = fmt.Sprintf("%v %v ", s, string(b[x][y]))
		}
		s = fmt.Sprintln(s)
	}
	return s
}

func (b Board) RowLen() int {
	return len(b)
}
func (b Board) ColLen() int {
	return len(b[0])
}
func (b Board) LastCoord() *Coord {
	return &Coord{x: b.RowLen() - 1, y: b.ColLen() - 1}
}
func (b Board) Value(c *Coord) byte {
	return b[c.x][c.y]
}

func (c *Coord) String() string {
	return fmt.Sprintf("x:%d,y:%d", c.x, c.y)
}
