package _0412

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	if n <= 0 {
		return nil
	}
	for i:=1;i<=n;i++{
		if i % 15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i % 3 == 0 {
			result = append(result, "Fizz")
		} else if i % 5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// 参考
func fizzBuzz1(n int) []string {
	result := make([]string, 0, n)
	var i int
	for i = 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i % 3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i % 5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, strconv.Itoa(i))
	}
	return result
}

func fizzBuzz2(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		newValue := ""

		if i%3 == 0 {
			newValue += "Fizz"
		}
		if i%5 == 0 {
			newValue += "Buzz"
		}

		if newValue == "" {
			newValue = strconv.Itoa(i)
		}

		result = append(result, newValue)
	}

	return result
}