package leetcode

func minCostClimbingStairs(cost []int) int {
	m := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	for i := 2; i < len(cost); i++ {
		cost[i] += m(cost[i-1], cost[i-2])
	}

	l := len(cost)
	return m(cost[l-1], cost[l-2])
}
