package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(is_valid_ip("1.2.3.4"))         //true
	fmt.Println(is_valid_ip("1.2.3.4.5"))       //false
	fmt.Println(is_valid_ip("123.45.67.89"))    //true
	fmt.Println(is_valid_ip("123.045.067.089")) //false

}

func is_valid_ip(ip string) bool {
	a := strings.Count(ip, ".")
	if a != 3 {
		return false
	}
	b := strings.SplitN(ip, ".", 4)
	var t2 = []int{}
	for _, i := range b {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}

	for _, v := range t2 {
		if v > 255 {
			return false
		}
		if v < 0 {
			return false
		}
	}
	return true
}
