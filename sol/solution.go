package sol

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	nIntervals := len(intervals)
	overlapStart, overlapEnd := newInterval[0], newInterval[1]
	hasInsert := false
	for pos := 0; pos < nIntervals; pos++ {
		if intervals[pos][1] < overlapStart {
			result = append(result, intervals[pos])
		} else if intervals[pos][0] > overlapEnd {
			if !hasInsert {
				result = append(result, []int{overlapStart, overlapEnd})
				hasInsert = true
			}
			result = append(result, intervals[pos])
		} else {
			if intervals[pos][0] <= overlapStart && overlapStart <= intervals[pos][1] {
				overlapStart = intervals[pos][0]
			}
			if intervals[pos][0] <= overlapEnd && overlapEnd <= intervals[pos][1] {
				overlapEnd = intervals[pos][1]
			}
		}
	}
	if !hasInsert {
		result = append(result, []int{overlapStart, overlapEnd})
		hasInsert = true
	}
	return result
}
