package main

func main() {
}

type Person struct {
	HttpStatus int
}

func fn() (name1 string, name2 string, name3 string, name4 string, name5 string, name6 string, name7 string, name8 string, name9 string, name10 string) {
	return "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"
}

func Add(a, b int) int {
	return a + b
}

func DealCards() (player1 []string, player2 []string) {
	player1 = append(player1, "name1")
	player2 = append(player2, "name2")

	return
}
