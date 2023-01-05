package main

import "testing"

func TestFindWrap(t *testing.T) {
	testCases := []struct {
		name       string
		currCol    int
		currRow    int
		nextCol    int
		nextRow    int
		currFacing string
		nextFacing string
	}{
		{
			name:       "A - T",
			currCol:    50,
			currRow:    53,
			nextCol:    49,
			nextRow:    53,
			currFacing: "T",
			nextFacing: "T",
		},
		{
			name:       "A - B",
			currCol:    99,
			currRow:    53,
			nextCol:    100,
			nextRow:    53,
			currFacing: "B",
			nextFacing: "B",
		},
		{
			name:       "A - R",
			currCol:    53,
			currRow:    99,
			nextCol:    49,
			nextRow:    103,
			currFacing: "R",
			nextFacing: "T",
		},
		{
			name:       "A - L",
			currCol:    53,
			currRow:    50,
			nextCol:    100,
			nextRow:    3,
			currFacing: "L",
			nextFacing: "B",
		},
		{
			name:       "B - T",
			currCol:    0,
			currRow:    53,
			nextCol:    153,
			nextRow:    0,
			currFacing: "T",
			nextFacing: "R",
		},
		{
			name:       "B - B",
			currCol:    49,
			currRow:    53,
			nextCol:    50,
			nextRow:    53,
			currFacing: "B",
			nextFacing: "B",
		},
		{
			name:       "B - R",
			currCol:    3,
			currRow:    99,
			nextCol:    3,
			nextRow:    100,
			currFacing: "R",
			nextFacing: "R",
		},
		{
			name:       "B - L",
			currCol:    3,
			currRow:    50,
			nextCol:    146,
			nextRow:    0,
			currFacing: "L",
			nextFacing: "R",
		},
		{
			name:       "C - T",
			currCol:    0,
			currRow:    103,
			nextCol:    199,
			nextRow:    3,
			currFacing: "T",
			nextFacing: "T",
		},
		{
			name:       "C - B",
			currCol:    49,
			currRow:    103,
			nextCol:    53,
			nextRow:    99,
			currFacing: "B",
			nextFacing: "L",
		},
		{
			name:       "C - R",
			currCol:    3,
			currRow:    149,
			nextCol:    146,
			nextRow:    99,
			currFacing: "R",
			nextFacing: "L",
		},
		{
			name:       "C - L",
			currCol:    3,
			currRow:    100,
			nextCol:    3,
			nextRow:    99,
			currFacing: "L",
			nextFacing: "L",
		},
		{
			name:       "D - T",
			currCol:    100,
			currRow:    53,
			nextCol:    99,
			nextRow:    53,
			currFacing: "T",
			nextFacing: "T",
		},
		{
			name:       "D - B",
			currCol:    149,
			currRow:    53,
			nextCol:    153,
			nextRow:    49,
			currFacing: "B",
			nextFacing: "L",
		},
		{
			name:       "D - R",
			currCol:    103,
			currRow:    99,
			nextCol:    46,
			nextRow:    149,
			currFacing: "R",
			nextFacing: "L",
		},
		{
			name:       "D - L",
			currCol:    103,
			currRow:    50,
			nextCol:    103,
			nextRow:    49,
			currFacing: "L",
			nextFacing: "L",
		},
		{
			name:       "E - T",
			currCol:    100,
			currRow:    3,
			nextCol:    53,
			nextRow:    50,
			currFacing: "T",
			nextFacing: "R",
		},
		{
			name:       "E - B",
			currCol:    149,
			currRow:    3,
			nextCol:    150,
			nextRow:    3,
			currFacing: "B",
			nextFacing: "B",
		},
		{
			name:       "E - R",
			currCol:    103,
			currRow:    49,
			nextCol:    103,
			nextRow:    50,
			currFacing: "R",
			nextFacing: "R",
		},
		{
			name:       "E - L",
			currCol:    103,
			currRow:    0,
			nextCol:    46,
			nextRow:    50,
			currFacing: "L",
			nextFacing: "R",
		},
		{
			name:       "F - T",
			currCol:    150,
			currRow:    3,
			nextCol:    149,
			nextRow:    3,
			currFacing: "T",
			nextFacing: "T",
		},
		{
			name:       "F - B",
			currCol:    199,
			currRow:    3,
			nextCol:    0,
			nextRow:    103,
			currFacing: "B",
			nextFacing: "B",
		},
		{
			name:       "F - R",
			currCol:    153,
			currRow:    49,
			nextCol:    149,
			nextRow:    53,
			currFacing: "R",
			nextFacing: "T",
		},
		{
			name:       "F - L",
			currCol:    153,
			currRow:    0,
			nextCol:    0,
			nextRow:    53,
			currFacing: "L",
			nextFacing: "B",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			nextFacing, nextCol, nextRow := findWrap(testCase.currCol, testCase.currRow, testCase.currFacing)
			if nextFacing != testCase.nextFacing {
				t.Fatalf("expect facing %s but got %s", testCase.nextFacing, nextFacing)
			}
			if nextCol != testCase.nextCol {
				t.Fatalf("expect column %d but got %d", testCase.nextCol, nextCol)
			}
			if nextRow != testCase.nextRow {
				t.Fatalf("expect row %d but got %d", testCase.nextRow, nextRow)
			}
		})
	}
}
