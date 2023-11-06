package main

func forlen1() {
	values := []string{
		"a", "b", "c", "d", "e", "f",
	}

	for i := 0; i < len(values); i++ {
	}
}

func forlen2() {
	values := []string{
		"a", "b", "c", "d", "e", "f",
	}

	l := len(values)
	for i := 0; i < l; i++ {
	}
}
