package main

import (
	//"bytes"
	//"encoding/json"
	//"file/filepath"
	"fmt"
	//"os"
	//"strings"
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	board := NewBoard(Board5x5)

	fmt.Println("")
	fmt.Println(board)

	words := findWordsInBoard(board)

	stdin := bufio.NewReader(os.Stdin)

	fmt.Println("Type a word to check and then hit <Enter>")
	fmt.Println("To finish, press the <Enter> key twice")

	var empty bool
	for {
		solution := &Solution{}
		line, err := stdin.ReadString('\n')
		if err != nil {
			log.Fatalln("Caught err on stdin:", err)
		}

		line = strings.TrimSuffix(line, "\n")
		line = strings.ToUpper(line)
		if line == "" {
			if empty {
				break
			}
			empty = true
		} else if len(line) <= 2 {
			fmt.Println("Word", line, "too short")
		} else {
			empty = false
			if boggleHasWord(board, line, solution) {
				fmt.Println("Word", line, "is in the board")
			} else {
				fmt.Println("Word", line, "not found")
			}
		}
	}

	fmt.Printf("Found %d Words\n\n", len(words))
	for _, word := range words {
		if len(word) > 2 {
			fmt.Println("Found word:", word)
		}
	}
}

func findWordsInBoard(board Board) []string {
	b, err := ioutil.ReadFile("US.dic")
	if err != nil {
		fmt.Println("Could not read file:", err)
		return nil
	}

	var words []string
	buf := bytes.NewBuffer(b)
	for {
		solution := &Solution{}
		line, err := buf.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err == nil {
			if line != "" {
				if boggleHasWord(board, line, solution) {
					words = append(words, line)
				}
			}
		} else {
			if err == io.EOF {
				if line != "" {
					if boggleHasWord(board, line, solution) {
						words = append(words, line)
					}
				}
			} else {
				fmt.Print("Caught error reading file lines:", err)
			}
			break
		}
	}
	return words
}

func boggleHasWord(board Board, word string, solution *Solution) bool {
	var searchCoords []*Coord

	boardSize := board.LastCoord()

	if solution.Len() > len(word) {
		return false
	} else if solution.Len() == 0 {
		searchCoords = getEntireBoard(boardSize)
	} else if solution.Len() == len(word) {
		return true
	} else {
		searchCoords = getAdjacents(solution.Last(), boardSize)
	}

	searchLetter := word[solution.Len()]
	for _, c := range searchCoords {
		if board.Value(c) == searchLetter {
			if solution.addCoord(c) {
				return boggleHasWord(board, word, solution)
			}
		}
	}

	return false
}

func getEntireBoard(max *Coord) []*Coord {
	var coords []*Coord

	for x := 0; x <= max.x; x++ {
		for y := 0; y <= max.y; y++ {
			coords = append(coords, &Coord{x: x, y: y})
		}
	}

	return coords
}

func getAdjacents(current, max *Coord) []*Coord {
	var adjacents []*Coord

	for x := current.x - 1; x <= current.x+1; x++ {
		for y := current.y - 1; y <= current.y+1; y++ {
			if x >= 0 && y >= 0 && x <= max.x && y <= max.y {
				if x != current.x || y != current.y {
					adjacents = append(adjacents, &Coord{x: x, y: y})
				}
			}
		}
	}

	return adjacents
}
