package main

import "fmt"

func main() {
	m := make(map[int]int, 0)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	for _, a := range m {
		fmt.Println(a)
	}
}
