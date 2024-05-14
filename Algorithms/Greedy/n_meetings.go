package main

import (
	"fmt"
	"sort"
)

// MaxMeetings provides max meeting given start and end time
func MaxMeetings(start []int, end []int, n int) int {
	count := 1
	meeting := make([][2]int, n)

	for i := 0; i < n; i++ {
		meeting[i][0] = start[i]
		meeting[i][1] = end[i]
	}

	//sorting
	sort.Slice(meeting, func(i, j int) bool {
		return meeting[i][1] < meeting[j][1]
	})

	endTime := meeting[0][1]
	for i := 1; i < n; i++ {
		if meeting[i][0] > endTime {
			count++
			endTime = meeting[i][1]
		}
	}

	return count
	
}

func main() {
	start := []int{1, 3, 0, 5, 8, 5}
	end := []int{2, 4, 6, 7, 9, 9}
	n := len(start)

	fmt.Println("Maximum number of meetings:", MaxMeetings(start, end, n))
}
