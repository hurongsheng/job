package main

import "fmt"

func main() {
	k := 8
	start := make([]QueenS, 0)
	for i := 0; i < k; i++ {
		start = append(start, append(QueenS{}, QueenPos{0, i}))
	}
	queens := Queens(start, k)
	for _, queen := range queens {
		strs := queen.Map(k)
		for _, str := range strs {
			fmt.Println(str)
		}
		fmt.Println("=====")
	}

}

type QueenPos struct {
	X int
	Y int
}
type QueenS []QueenPos

func (qs *QueenS) Check(x, y int) bool {
	if len(*qs) == 0 {
		return true
	}
	for _, q := range *qs {
		if q.X == x || q.Y == y || (q.Y-q.X) == (y-x) || (q.X+q.Y) == (x+y) {
			return false
		}

	}
	return true
}
func (qs *QueenS) Map(k int) []string {
	xxx := string(X(k))
	ret := make([]string, k, k)
	for _, q := range *qs {
		xx := []byte(xxx)
		xx[q.Y] = 'Q'
		ret[q.X] = string(xx)
	}
	return ret
}
func X(k int) []byte {
	b := make([]byte, 0, k)
	for i := 0; i < k; i++ {
		b = append(b, '.')
	}
	return b
}

func Queens(queens []QueenS, k int) []QueenS {
	newQ := make([]QueenS, 0)
	for _, queen := range queens {
		if len(queen) == k {
			return queens
		}
		x := len(queen)
		for y := 0; y < k; y++ {
			if queen.Check(x, y) {
				newQ = append(newQ, append(queen, QueenPos{x, y}))
			}
		}
	}
	return Queens(newQ, k)
}
