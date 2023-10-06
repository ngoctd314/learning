package main

func flipAndInvertImage(image [][]int) [][]int {
	rs := make([][]int, 0, len(image))
	for _, v := range image {
		i, j := 0, len(v)-1
		for i < j {
			v[i], v[j] = v[j]^1, v[i]^1
			i++
			j--
		}
		rs = append(rs, v)
	}

	return rs
}
