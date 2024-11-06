package main

import (
	"fmt"
	"math"
)

/*
	Code bat solutions warmup1 java questions
*/

func main() {
	fmt.Println(sleepIn(true, false))
	fmt.Println(monkeyTrouble(false, false))
	fmt.Println(sumDouble(2, 2))
	fmt.Println(diff21(22))
	fmt.Println(parrotTrouble(true, 7))
	fmt.Println(makes10(9, 10))
	fmt.Println(nearHundred(90))
	fmt.Println(posNeg(-1, -1, true))
	fmt.Println(notString("not"))
}

func sleepIn(weekdays, vacation bool) bool {
	return !weekdays || vacation
}

func monkeyTrouble(aSmile, bSmile bool) bool {
	return aSmile == bSmile
}

func sumDouble(a, b int) int {
	if a == b {
		return (a + b) * 2
	}
	return a + b
}

func diff21(n int) int {
	if n <= 21 {
		return 21 - n
	}

	return (n - 21) * 2
}

func parrotTrouble(talking bool, hour int) bool {
	return talking && (hour < 7 || hour > 20)
}

func makes10(a, b int) bool {
	if a+b == 10 {
		return true
	}

	if a == 10 || b == 10 {
		return true
	}

	return false
}

func nearHundred(n int) bool {
	fmt.Println("Math abs -->", math.Abs(float64(n)))
	if n >= 100-10 && n <= 100+10 {
		return true
	}
	if n >= 200-10 && n <= 200+10 {
		return true
	}
	return false
}

func posNeg(a, b int, negative bool) bool {
	if negative {
		return a < 0 && b < 0
	}
	return (a < 0 && b > 0) || (a > 0 && b < 0)
}

func notString(str string) string {
	if len(str) >= 3 && str[0:3] == "not" {
		return str
	}

	return "not " + str
}
