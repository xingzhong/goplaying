package main

import (
	"fmt"
)

type Test struct {
	X int
	Y int
}

func (test *Test) String() string {
	return fmt.Sprintf("<%d, %d>", test.X, test.Y)
}

func main() {
	fmt.Println("[Example] map value with reference")

	originalArray := make([]Test, 10)
	hashMap := make(map[int]Test)
	hashRefMap := make(map[int]*Test)
	for i := 0; i < 10; i++ {
		originalArray[i] = Test{i, -i}
	}
	fmt.Printf("%v\n", originalArray)
	for i, _ := range originalArray {
		hashMap[i] = originalArray[i]
		hashRefMap[i] = &originalArray[i]
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d = {%v %v}\n", i, hashMap[i], hashRefMap[i])
	}
	fmt.Println("change the underlying data")
	for i := 0; i < 10; i++ {
		originalArray[i] = Test{2 * i, -i * 2}
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d = {%v %v}\n", i, hashMap[i], hashRefMap[i])
	}
	fmt.Println("only hashRefMap data changed")
}
