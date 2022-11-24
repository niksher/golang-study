package main

import (
	"fmt"
	"golang-study/lesson"
)

func main() {
	slice := []int{4, 3, 3, 3, 4, 6, 6, 6, 8}
	out := lesson.Lesson1(slice, 1)

	fmt.Println("lesson 1: ", out)

	slice = []int{4, 3, 3, 3, 4, 3, 6, 6, 8}
	outInt := lesson.Lesson2(slice)

	fmt.Println("lesson 2: ", outInt)

	slice = []int{3, 4, 5, 6, 8, 7, 9}
	outInt = lesson.Lesson3(slice)

	fmt.Println("lesson 3: ", outInt)

	slice = []int{4, 3, 6, 5, 8}
	outInt = lesson.Lesson4(slice)

	fmt.Println("lesson 4: ", outInt)
}
