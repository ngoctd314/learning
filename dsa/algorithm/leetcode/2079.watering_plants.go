package leetcode

func wateringPlants(plants []int, capacity int) int {
	var rs int
	cpCap := capacity
	for i := 0; i < len(plants); i++ {
		if capacity < plants[i] {
			capacity = cpCap
			rs += i * 2
		}
		capacity -= plants[i]
	}

	return rs + len(plants)
}
