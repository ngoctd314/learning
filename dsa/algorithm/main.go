package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	n := 6
	fmt.Println(n * (n + 1) / 2)
}

func cnt(n int) int {
	rs := 0
	for i := 0; i <= n; i++ {
		str := fmt.Sprint(i)
		for _, v := range str {
			if v == '1' {
				rs++
			}
		}
	}
	return rs

	// 1 + 0*1 + 1 : 10
	// 1 + 1*10 + 10 : 100
	// 1 + 2*100 + 100: 1000
	// 1 + 3*1000 + 1000: 10000
}

func iter(n int) int {
	str := fmt.Sprint(n)
	calc := func(n int, lenN int) int {
		return 1 + (n/10)*(lenN-1)
	}

	rs := 0
	trackingLen := len(str)
	for _, v := range str {
		firstNum, _ := strconv.Atoi(string(v))
		rs += calc(int(math.Pow(10, float64(trackingLen)))*firstNum, trackingLen)

		trackingLen--
	}
	return rs
}

func quickCalc(n int) int {
	if n < 1 {
		return 0
	}
	if n < 10 {
		return 1
	}

	str := fmt.Sprint(n)
	firstNum, _ := strconv.Atoi(string(str[0]))
	remainNum, _ := strconv.Atoi(string(str[1:]))
	powFloat := math.Pow(10, float64(len(str)-1))
	pow := int(powFloat)

	if firstNum == 1 {
		return calc(pow) + remainNum + quickCalc(remainNum)
	}

	return calc(pow)*firstNum + pow - firstNum + quickCalc(remainNum)
}

func calc(n int) int {
	return 1 + (n/10)*(len(fmt.Sprint(n))-1)
}

func sum(n int) int {
	return (1 + n) * (n / 2)
}
