package main

import "fmt"

/*
	Code bat solutions warmup2 java questions
*/

func main() {
	fmt.Println("String times -> ", stringTimes("Hi", 2))
	fmt.Println("Front times -> ", frontTimes("Ab", 4))
	fmt.Println("Count XX -> ", countXX("xxxx"))
	fmt.Println("Double X -> ", doubleX("x"))
	fmt.Println("String bits -> ", stringBits("Heeololeo"))
	fmt.Println("String Splosion -> ", stringSplosion("abc"))
	fmt.Println("Last 2 -> ", last2("xaxxaxaxx"))
	slicesCount9 := []int{1, 9, 9, 3, 9}
	fmt.Println("Array Count 9 -> ", arrayCount9(slicesCount9...))
	slicesFront9 := []int{1, 2, 9, 3, 4}
	fmt.Println("Slices Front 9 -> ", arrayFront9(slicesFront9...))
	slicesArray123 := []int{1, 1, 2, 1, 2, 3}
	fmt.Println("Array 123 -> ", array123(slicesArray123...))
}

func stringTimes(str string, n int) string {
	var temp string
	for i := 0; i < n; i++ {
		temp += str
	}
	return temp
}

func frontTimes(str string, n int) string {
	var temp string
	for i := 0; i < n; i++ {
		if len(str) > 3 {
			temp += str[0:3]
		} else {
			temp += str
		}
	}
	return temp
}

func countXX(str string) int {
	var count int
	for i := 0; i <= len(str); i++ {
		if i+1 == len(str) {
			break
		}
		if str[i:i+2] == "xx" {
			count += 1
		}
	}
	return count
}

func doubleX(str string) bool {
	var flag = false
	for i := 0; i <= len(str); i++ {
		if len(str) == 0 || i+1 == len(str) {
			break
		}
		if str[i:i+2] == "xx" {
			flag = true
		} else {
			break
		}
	}
	return flag
}

func stringBits(str string) string {
	var temp string
	for i := 0; i < len(str); i++ {
		if i%2 == 0 {
			temp += str[i : i+1]
		}
	}
	return temp
}

func stringSplosion(str string) string {
	var temp string
	for i := 0; i <= len(str); i++ {
		if i == len(str) {
			break
		}
		temp += str[0 : i+1]
	}
	return temp
}

func last2(str string) int {
	var count int
	for i := 0; i < len(str); i++ {
		if i+1 == len(str) {
			break
		}
		if str[i:i+2] == "xx" {
			count += 1
			i++
		}
	}

	return count
}

func arrayCount9(nums ...int) int {
	var count int
	for _, num := range nums {
		if num == 9 {
			count++
		}
	}
	return count
}

func arrayFront9(nums ...int) bool {
	var flag bool
	for i, num := range nums {
		if i <= 3 && num == 9 {
			flag = true
		}
	}
	return flag
}

func array123(nums ...int) bool {
	var flag bool
	for i := 0; i < len(nums); i++ {
		if i+2 == len(nums) || i+1 == len(nums) {
			break
		}
		if nums[i] == 1 && nums[i+1] == 2 && nums[i+2] == 3 {
			flag = true
		}
	}
	return flag
}
