package _01

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return result
}

func fizzBuzz1(n int) []string {
	var result = make([]string, n)
	for i := 1; i <= n; i++ {
		t := ""
		if i%3 == 0 {
			t += "Fizz"
		}

		if i%5 == 0 {
			t += "Buzz"
		}

		if t == "" {
			t = strconv.Itoa(i)
		}
		result[i-1] = t
	}
	return result
}
