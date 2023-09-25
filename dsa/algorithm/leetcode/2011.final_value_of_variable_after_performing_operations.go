package main

func finalValueAfterOperations(operations []string) int {
	r := 0
	for _, v := range operations {
		if v[0] == '-' || v[2] == '-' {
			r--
		} else {
			r++
		}
	}
	return r
}
