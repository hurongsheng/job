package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	line := make([]int, n)
	for i, _ := range line {
		fmt.Scanf("%d", &line[i])
	}
	//_, _ = fmt.Scanln(&n)
	fmt.Printf("aaa %v", line)
	return
	sortedNums := sortStringInt(line, n)
	need := 0
	for n >= 3 {
		if sortedNums[n-1]-sortedNums[n-2] > 10 {
			sortedNums[n] = sortedNums[n-1]
			n++
			need++
			continue
		}
		if sortedNums[n-2]-sortedNums[n-3] > 10 {
			n = n - 2
			need++
			continue
		}
		n = n - 3
	}
	if len(sortedNums) > 0 {
		need = need + 3 - len(sortedNums)
	}
	fmt.Println(need)

}
func sortStringInt(strs []int, n int) []int {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if strs[i] > strs[j] {
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}
	return strs
}
