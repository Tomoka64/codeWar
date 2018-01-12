package main

import "fmt"

func century(year int) int {
	// Finish this :)
	a := year / 100
	b := year % 100
	if b > 0 {
		a = a + 1
		return a
	}
	return a
}

func main() {
	fmt.Println(century(2019))
}

//other answer
// func century(year int) int {
//  return int(((year - 1) / 100)) + 1
// }

// func century(year int) int {
//  result := int(year/100)
//  if year % 100 != 0 {
//    result += 1
//  }
//  return result
// }
