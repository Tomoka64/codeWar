package main

import "fmt"

func SumEvenFibonacci(limit int) int {
	x, y := 0, 1
	for i := 0; ; i++ {
		c := make(chan int)
		x, y = y, x+y
		if x%2 == 0 {
			c <- x
			if x <= limit {
				for v := range c {
					v += v
					return v
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(SumEvenFibonacci(8))
}
