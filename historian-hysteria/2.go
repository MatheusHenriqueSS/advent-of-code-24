package main

import (
	"fmt"
)

func main() {

	var a, b int

	var vectorA []int
	mp := make(map[int]int)

	for {
		_, err := fmt.Scanf("%d %d", &a, &b)

		if err != nil {
			break
		}

		vectorA = append(vectorA, a)
		mp[b]++
	}

	ans := 0

	for _, value := range vectorA {
		ans += value * mp[value]
	}

	fmt.Printf("result %d", ans)

}
