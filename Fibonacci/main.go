package main

import "fmt"

func SumEvenFibonacci(limit int) int {
	sum, a, b := 2, 1, 2
	for b <= limit {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	return sum
}

func main() {
	fmt.Println(SumEvenFibonacci(111111))
}
