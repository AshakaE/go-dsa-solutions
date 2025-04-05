package samples

import (
	"fmt"
	"github.com/AshakaE/go-dsa-solutions/algorithms"
)

func DiagonalDifferenceTest() {
    testCases := []struct {
		name     string
		input    [][]int32
		expected int32
	}{
		{
			name: "3x3 matrix",
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9},
			},
			expected: 2,
		},
		{
			name: "2x2 matrix",
			input: [][]int32{
				{11, 2},
				{4, 5},
			},
			expected: 10,
		},
		{
			name: "4x4 matrix",
			input: [][]int32{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			expected: 8,
		},
		{
			name:     "1x1 matrix",
			input:    [][]int32{{5}},
			expected: 0,
		},
		{
			name: "Negative numbers",
			input: [][]int32{
				{-1, -2, -3},
				{-4, -5, -6},
				{-9, -8, -9},
			},
			expected: 2,
		},
	}

	for i, tc := range testCases {
        result := algorithms.DiagonalDifference(tc.input)
        status := "FAILED"
        if result == tc.expected {
            status = "PASSED"
        }
        fmt.Printf("Test case %d %s - Input: %v, Result: %d, Expected: %d\n",
            i+1,
            status,
            tc.input,
            result,
            tc.expected)
    }
}