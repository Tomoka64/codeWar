package main

import (
	"fmt"
)

//
// func ListSquared(m, n int) int {
// 	// your code
// 	var a []int
// 	a = getDivisor(n)
// 	sum := 0
// 	for _, v := range a {
// 		b := v ^ 2
// 		sum += b
// 	}
// 	return sum
// }

func getDivisor(initNum int, targetNum int) []int {
	var retFrontSlice []int
	var retBackSlice []int
	var loopMax int = targetNum
	var i int = 2
	for initNum <= targetNum {
		retFrontSlice = append(retFrontSlice, initNum)
		if targetNum > initNum {
			retBackSlice = append(retBackSlice, targetNum)
		}
		for ; i < loopMax; i++ {
			if targetNum%i == 0 {
				loopMax = targetNum / i
				retFrontSlice = append(retFrontSlice, i)
				if i != loopMax {
					retBackSlice = append(retBackSlice, loopMax)
				}
			}
		}
		for j := len(retBackSlice) - 1; j >= 0; j-- {
			retFrontSlice = append(retFrontSlice, retBackSlice[j])
		}
		return retFrontSlice
	}
	return nil
}

func nijou(a []int) []int {
	var b []int
	for _, v := range a {
		ni := v * v
		b = append(b, ni)
	}
	return b
}

func sum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	var a []int
	a = getDivisor(1, 250)
	fmt.Println(a)
	a = nijou(a)
	fmt.Println(a)
	b := sum(a)
	fmt.Println(b)
}
