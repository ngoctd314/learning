package leetcode

// TODO: optimize
func countPoints(points [][]int, queries [][]int) []int {
	rs := make([]int, len(queries))
	fn := func(x, y int, r int, x1, y1 int) bool {
		p1 := x - x1
		p2 := y - y1
		return p1*p1+p2*p2 <= r*r
	}
	for i, query := range queries {
		tmp := 0
		for _, point := range points {
			if fn(query[0], query[1], query[2], point[0], point[1]) {
				tmp++
			}
		}
		rs[i] = tmp
	}

	return rs
}
