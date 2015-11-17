package main

import (
	//"bytes"
	//"encoding/json"
	//"file/filepath"
	//"fmt"
	//"os"
	//"strings"
	//"fmt"
	"testing"
)

var board2x2 = [][]byte{{'b', 'a'}, {'r', 't'}}
var board3x3 = [][]byte{{'b', 'a', 'e'}, {'r', 't', 'm'}, {'z', 'z', 'z'}}

/*
	0	1	2	3
	1	-	-	-
	2	-	-	-
	3	-	-	-
*/
var adjacent_testcases = []struct {
	start     *Coord
	max       *Coord
	expecting []*Coord
}{
	{Row0Col0, BoardMax2x2, []*Coord{Row0Col1, Row1Col0, Row1Col1}},
	{Row0Col0, BoardMax3x3, []*Coord{Row0Col1, Row1Col0, Row1Col1}},
	{Row0Col0, BoardMax4x4, []*Coord{Row0Col1, Row1Col0, Row1Col1}},

	{Row0Col1, BoardMax2x2, []*Coord{Row0Col0, Row1Col0, Row1Col1}},
	{Row0Col1, BoardMax3x3, []*Coord{Row0Col0, Row0Col2, Row1Col0, Row1Col1, Row1Col2}},
	{Row0Col1, BoardMax4x4, []*Coord{Row0Col0, Row0Col2, Row1Col0, Row1Col1, Row1Col2}},

	{Row0Col2, BoardMax3x3, []*Coord{Row0Col1, Row1Col1, Row1Col2}},
	{Row0Col2, BoardMax4x4, []*Coord{Row0Col1, Row0Col3, Row1Col1, Row1Col2, Row1Col3}},

	{Row0Col3, BoardMax4x4, []*Coord{Row0Col2, Row1Col2, Row1Col3}},

	{Row1Col0, BoardMax2x2, []*Coord{Row0Col0, Row0Col1, Row1Col1}},
	{Row1Col0, BoardMax3x3, []*Coord{Row0Col0, Row0Col1, Row1Col1, Row2Col0, Row2Col1}},
	{Row1Col0, BoardMax4x4, []*Coord{Row0Col0, Row0Col1, Row1Col1, Row2Col0, Row2Col1}},
}

var board_testcases = []struct {
	board        [][]byte
	expectations map[string]bool
}{
	{board2x2, map[string]bool{"bar": true, "bart": true, "barb": false, "back": false, "barth": false, "rath": false, "rat": true}},
	{board3x3, map[string]bool{"bar": true, "bart": true, "barb": false, "back": false, "barth": false, "rath": false, "rat": true, "tram": true, "zz": true, "zzz": true, "zzzz": false}},
}

func TestBoards(t *testing.T) {
	for _, tc := range board_testcases {
		for test, expected := range tc.expectations {
			actual := boggleHasWord(tc.board, test, &Solution{})
			if actual != expected {
				t.Logf("Expected test [%v] to have result %v", test, expected)
				t.Fail()
			} else {
				if actual {
					t.Logf("Success: Found test word [%v]\n", test)
				} else {
					t.Logf("Success: Did not find test word [%v]\n", test)
				}
			}
		}
	}
}
func TestAdjacents(t *testing.T) {

	for _, tc := range adjacent_testcases {
		actual := getAdjacents(tc.start, tc.max)
		if !match(actual, tc.expecting) {
			t.Logf("Starting at coord %+v and expecting %+v, but found %+v", tc.start, tc.expecting, actual)
			t.Fail()
		}
	}
}

func match(actual, expectation []*Coord) bool {
	if len(actual) != len(expectation) {
		return false
	}

outer:
	for _, a := range actual {
		for _, e := range expectation {
			if e.x == a.x && e.y == a.y {
				continue outer
			}
		}
		return false
	}
	return true
}

//func boggleHasWord(board [][]byte, word string, solution *Solution) bool {

//func getEntireBoard(max *Coord) []*Coord {

//func getAdjacents(current, max *Coord) []*Coord {

var (
	Row0Col0 *Coord = &Coord{x: 0, y: 0}
	Row0Col1        = &Coord{x: 0, y: 1}
	Row0Col2        = &Coord{x: 0, y: 2}
	Row0Col3        = &Coord{x: 0, y: 3}

	Row1Col0 = &Coord{x: 1, y: 0}
	Row1Col1 = &Coord{x: 1, y: 1}
	Row1Col2 = &Coord{x: 1, y: 2}
	Row1Col3 = &Coord{x: 1, y: 3}

	Row2Col0 = &Coord{x: 2, y: 0}
	Row2Col1 = &Coord{x: 2, y: 1}
	Row2Col2 = &Coord{x: 2, y: 2}
	Row2Col3 = &Coord{x: 2, y: 3}

	Row3Col0 = &Coord{x: 3, y: 0}
	Row3Col1 = &Coord{x: 3, y: 1}
	Row3Col2 = &Coord{x: 3, y: 2}
	Row3Col3 = &Coord{x: 3, y: 3}

	BoardMax2x2 = Row1Col1
	BoardMax3x3 = Row2Col2
	BoardMax4x4 = Row3Col3
)
