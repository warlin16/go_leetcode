package easy

import "sort"

var cost = [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(a, b int) bool { return costs[a][0]-costs[a][1] < costs[b][0]-costs[b][1] })
	result := 0
	for i := 0; i < len(costs); i++ {
		if i < len(costs)/2 {
			result += costs[i][0]
		} else {
			result += costs[i][1]
		}
	}
	return result
}
