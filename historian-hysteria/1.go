package main

import (
	"fmt"
	"math"
	"sort"
)

func notmain() {

	var a, b int

	var vectorA []int
	var vectorB []int

	for {
		_, err := fmt.Scanf("%d %d", &a, &b)

		if err != nil {
			break
		}

		vectorA = append(vectorA, a)
		vectorB = append(vectorB, b)
	}

	sort.Ints(vectorA)
	sort.Ints(vectorB)

	ans := 0

	for i := range vectorA {
		ans += int(math.Abs(float64(vectorA[i] - vectorB[i])))

	}

	fmt.Printf("result %d", ans)

}
