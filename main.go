package main

import (
	"fmt"
	"sort"
)

const (
	lowValue  = 0
	highValue = 1
)

func main() {
	example := [][]int{{2, 6}, {9, 12}, {8, 9}, {18, 21}, {4, 7}, {10, 11}}

	uncovered := uncoveredIntervals(example)
	fmt.Printf("uncovered: %v\n", uncovered)
}

// overlapIntervals finds overlapping ranges in a slice of ranges
func overlapIntervals(intervals [][]int) [][]int {
	// Sort the ranges by the low value.
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][lowValue] < intervals[j][lowValue] })

	var overlap [][]int
	for i := 0; i < len(intervals); i++ {
		low := intervals[i][lowValue]
		high := intervals[i][highValue]

		// This should take care of any ranges that is already covered from the previous add
		// eg:  overlap already has {{2,15}} and we are evaluating {4, 5}
		if overlap != nil {
			lastOverlap := overlap[len(overlap)-1]
			if high <= lastOverlap[highValue] {
				continue
			}
		}

		// Iterate through the rest of the values until you find a value that does not overlap.
		for j := i + 1; j < len(intervals); j++ {
			// make sure the next item low value is in range
			if intervals[j][lowValue] <= high {
				// make sure the next item high value is higher than what we already have for high
				if intervals[j][highValue] > high {
					high = intervals[j][highValue]
				}
			} else {
				// no need to process the rest since we are working on a sorted array
				break
			}
		}
		overlap = append(overlap, []int{low, high})
	}

	return overlap
}

// uncoveredIntervals finds uncovered ranges in a slice of ranges
func uncoveredIntervals(intervals [][]int) [][]int {
	overlap := overlapIntervals(intervals)

	var uncovered [][]int
	for i := 1; i < len(overlap); i++ {
		uncovered = append(uncovered, []int{overlap[i-1][1], overlap[i][0]})
	}

	return uncovered
}
