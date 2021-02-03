package main

import (
	"fmt"
)

type s struct {
	Key  string
	Key2 int
}

func main() {
	CompareStruct()
}
func CompareStruct() {
	int1 := 6
	int2 := 6
	s1 := s{
		Key:  "1",
		Key2: int1,
	}
	s2 := s{
		Key:  "1",
		Key2: int2,
	}
	fmt.Printf("%+v == %+v? %v\n", s1, s2, s1 == s2)
	fmt.Printf("%+v == %+v? %v\n", s1, s2, &s1 == &s2)
}
