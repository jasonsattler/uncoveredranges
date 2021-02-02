package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOverlapIntervals(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{2, 6}, {9, 12}, {8, 9}, {18, 21}, {4, 7}, {10, 11}},
			expected: [][]int{{2, 7}, {8, 12}, {18, 21}},
		},
		{
			input:    [][]int{{2, 3}, {9, 11}, {8, 9}, {15, 21}, {4, 7}, {5, 11}},
			expected: [][]int{{2, 3}, {4, 11}, {15, 21}},
		},
		{
			input:    [][]int{{9, 12}, {8, 9}, {18, 21}, {4, 7}, {10, 11}, {2, 15}},
			expected: [][]int{{2, 15}, {18, 21}},
		},
		{
			input:    [][]int{{9, 12}},
			expected: [][]int{{9, 12}},
		},
	}

	for _, test := range tests {
		overlap := overlapIntervals(test.input)
		require.Equal(t, test.expected, overlap)
	}
}

func TestUncoveredIntervals(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{2, 6}, {9, 12}, {8, 9}, {18, 21}, {4, 7}, {10, 11}},
			expected: [][]int{{7, 8}, {12, 18}},
		},
		{
			input:    [][]int{{2, 3}, {9, 11}, {8, 9}, {15, 21}, {4, 7}, {5, 11}},
			expected: [][]int{{3, 4}, {11, 15}},
		},
		{
			input:    [][]int{{9, 12}, {8, 9}, {18, 21}, {4, 7}, {10, 11}, {2, 15}},
			expected: [][]int{{15, 18}},
		},
		{
			input:    [][]int{{9, 12}},
			expected: nil,
		},
	}

	for _, test := range tests {
		uncovered := uncoveredIntervals(test.input)
		require.Equal(t, test.expected, uncovered)
	}
}
