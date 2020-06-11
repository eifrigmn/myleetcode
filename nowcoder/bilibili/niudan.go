package main

import "fmt"

func main(){
	var count int
	cnt, err := fmt.Scanf("%d", &count)
	if cnt != 1 ||err != nil {
		return
	}

	var result []int
	for count != 0 {
		if count %2 == 0 {
			result = append(result, 3)
			count = (count-2)/2
		} else {
			result = append(result, 2)
			count = (count-1)/2
		}
	}
	for i:=len(result)-1; i>=0;i--{
		fmt.Printf("%d", result[i])
	}
	fmt.Println()
}

