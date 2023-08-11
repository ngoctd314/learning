package main

import "fmt"

// ConvStringNumbersToByteNumbers convert string numbers -> byte numbers
func ConvStringNumbersToByteNumbers(str string) []byte {
	var rs []byte
	for _, v := range []rune(str) {
		rs = append(rs, byte(v-48))
	}

	return rs
}

// ConvByteNumbersToStringNumbers convert byte numbers -> string numbers
func ConvByteNumbersToStringNumbers(b []byte) string {
	rs := ""
	for _, v := range b {
		rs += fmt.Sprint(v)
	}

	return rs
}
