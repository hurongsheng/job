package main

import "fmt"

func main() {

	arr := make([]int, 0, 12)
	i := 12
	arr = append(arr, i)
	fmt.Printf("%+v\n", &arr[0])
	arr = appendSlice(arr)
	fmt.Printf("%+v\n", &arr[0])
	appendSlice2(arr)
	fmt.Printf("%+v\n", &arr[0])

	testSlice()
	pase_student()
}

func appendSlice(arr []int) []int {
	i := 11
	arr = append(arr, i)
	return arr
}

func appendSlice2(arr []int) {
	i := 10
	arr = append(arr, i)
}

func testSlice() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5] //2,3,4
	s1[2] = 20
	fmt.Println(s1)
	s1 = append(s1, 30)
	fmt.Println(s1)
	fmt.Println(slice)

	fmt.Println("s2")
	s2 := s1[2:6:7] //4,5,6,7
	fmt.Println(s2)
	s2 = append(s2, 100)
	s2 = append(s2, 200)
	s1[2] = 20
	fmt.Println(s2)
	fmt.Println(slice)
}

type student2 struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student2)
	stus := []student2{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Printf("%+v", m)

}
